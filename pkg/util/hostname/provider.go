package hostname

import (
	"context"
	"expvar"
	"fmt"

	"opsAgent/pkg/util/cache"
	"opsAgent/pkg/util/log"
)

const (
	configProvider  = "configuration"
	fargateProvider = "fargate"
)

var (
	hostnameExpvars  = expvar.NewMap("hostname")
	hostnameProvider = expvar.String{}
	hostnameErrors   = expvar.Map{}
)

func init() {
	hostnameErrors.Init()
	hostnameExpvars.Set("provider", &hostnameProvider)
	hostnameExpvars.Set("errors", &hostnameErrors)
}

type providerCb func(ctx context.Context, currentHostname string) (string, error)

type provider struct {
	name string
	cb   providerCb

	stopIfSuccessful bool

	expvarName string
}

var providerCatalog = []provider{
	{
		name:             configProvider,
		cb:               fromConfig,
		stopIfSuccessful: true,
		expvarName:       "'hostname' configuration/environment",
	},
	{
		name:             "hostnameFile",
		cb:               fromHostnameFile,
		stopIfSuccessful: true,
		expvarName:       "'hostname_file' configuration/environment",
	},
	{
		name:             "fqdn",
		cb:               fromFQDN,
		stopIfSuccessful: false,
		expvarName:       "fqdn",
	},
	{
		name:             "os",
		cb:               fromOS,
		stopIfSuccessful: false,
		expvarName:       "os",
	},
}

func (h Data) FromConfiguration() bool {
	return h.Provider == configProvider
}

func saveHostname(cacheHostnameKey string, hostname string, providerName string) Data {
	data := Data{
		Hostname: hostname,
		Provider: providerName,
	}

	cache.Cache.Set(cacheHostnameKey, data, cache.NoExpiration)

	return data
}

func GetWithProvider(ctx context.Context) (Data, error) {

	cacheHostnameKey := cache.BuildAgentKey("hostname")

	if cacheHostname, found := cache.Cache.Get(cacheHostnameKey); found {
		return cacheHostname.(Data), nil
	}

	var err error
	var hostname string
	var providerName string

	for _, p := range providerCatalog {
		log.Debugf("trying to get hostname from '%s' provider", p.name)

		detectedHostname, err := p.cb(ctx, hostname)
		if err != nil {
			expErr := new(expvar.String)
			expErr.Set(err.Error())
			hostnameErrors.Set(p.expvarName, expErr)
			log.Debugf("unable to get the hostname from '%s' provider: %s", p.name, err)
			continue
		}

		log.Debugf("hostname provider '%s' successfully found hostname '%s'", p.name, detectedHostname)
		hostname = detectedHostname
		providerName = p.name

		if p.stopIfSuccessful {
			log.Debugf("hostname provider '%s' succeeded, stoping here with hostname '%s'", p.name, detectedHostname)
			return saveHostname(cacheHostnameKey, hostname, p.name), nil
		}
	}

	warnAboutFQDN(ctx, hostname)

	if hostname != "" {
		return saveHostname(cacheHostnameKey, hostname, providerName), nil
	}

	err = fmt.Errorf("unable to reliably determine the host name. You can define one in the agent config file or in your hosts file")
	expErr := new(expvar.String)
	expErr.Set(err.Error())
	hostnameErrors.Set("all", expErr)
	return Data{}, err
}

func Get(ctx context.Context) (string, error) {
	data, err := GetWithProvider(ctx)
	return data.Hostname, err
}
