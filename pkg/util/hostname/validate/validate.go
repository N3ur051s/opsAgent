package validate

import (
	"fmt"
	"regexp"
	"strings"

	"opsAgent/pkg/util/log"
)

const maxLength = 255

var (
	validHostnameRfc1123 = regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
	localhostIdentifiers = []string{
		"localhost",
		"localhost.localdomain",
		"localhost6.localdomain6",
		"ip6-localhost",
	}
)

func ValidHostname(hostname string) error {
	if hostname == "" {
		return fmt.Errorf("hostname is empty")
	} else if isLocal(hostname) {
		return fmt.Errorf("%s is a local hostname", hostname)
	} else if len(hostname) > maxLength {
		log.Errorf("ValidHostname: name exceeded the maximum length of %d characters", maxLength)
		return fmt.Errorf("name exceeded the maximum length of %d characters", maxLength)
	} else if !validHostnameRfc1123.MatchString(hostname) {
		log.Errorf("ValidHostname: %s is not RFC1123 compliant", hostname)
		return fmt.Errorf("%s is not RFC1123 compliant", hostname)
	}
	return nil
}

func isLocal(name string) bool {
	name = strings.ToLower(name)
	for _, val := range localhostIdentifiers {
		if val == name {
			return true
		}
	}
	return false
}
