package hostname

import (
	"context"
	"os"

	"simpleagent/pkg/util/log"
)

func warnIfNotCanonicalHostname(ctx context.Context, hostname string) {
	log.Warnf(
		"Hostname '%s' defined in configuration will not be used as the in-app hostname. ",
		hostname,
	)

}

func warnAboutFQDN(ctx context.Context, hostname string) {
	fqdn, _ := fromFQDN(ctx, "")
	if fqdn == "" {
		return
	}

	h, err := os.Hostname()
	if err != nil {
		return
	}

	if hostname == h && h != fqdn {
		log.Warnf("DEPRECATION NOTICE: The agent resolved your hostname as '%s'. However in a future version, it will be resolved as '%s' by default. ", h, fqdn)
	}
}
