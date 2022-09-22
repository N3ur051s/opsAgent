package hostname

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	. "simpleagent/conf"

	"simpleagent/pkg/util/hostname/validate"
)

var (
	osHostname   = os.Hostname
	fqdnHostname = getSystemFQDN
)

type Data struct {
	Hostname string
	Provider string
}

func fromConfig(ctx context.Context, _ string) (string, error) {
	configName := Conf.Server.Hostname
	err := validate.ValidHostname(configName)
	if err != nil {
		return "", err
	}

	return configName, nil
}

func fromHostnameFile(ctx context.Context, _ string) (string, error) {
	// Try `hostname_file` config option next
	hostnameFilepath := "/etc/hostname"
	if hostnameFilepath == "" {
		return "", fmt.Errorf("'hostname_file' configuration is not enabled")
	}

	fileContent, err := ioutil.ReadFile(hostnameFilepath)
	if err != nil {
		return "", fmt.Errorf("Could not read hostname from %s: %v", hostnameFilepath, err)
	}

	hostname := strings.TrimSpace(string(fileContent))

	err = validate.ValidHostname(hostname)
	if err != nil {
		return "", err
	}
	warnIfNotCanonicalHostname(ctx, hostname)
	return hostname, nil
}

func fromOS(ctx context.Context, currentHostname string) (string, error) {
	if currentHostname == "" {
		return osHostname()
	}
	return "", fmt.Errorf("Skipping OS hostname as a previous provider found a valid hostname")

}

func fromFQDN(ctx context.Context, _ string) (string, error) {
	fqdn, err := fqdnHostname()
	if err == nil {
		return fqdn, nil
	}
	return "", fmt.Errorf("Unable to get FQDN from system: %s", err)
}
