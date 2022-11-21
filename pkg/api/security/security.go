package security

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"opsAgent/pkg/util/filesystem"
	"opsAgent/pkg/util/log"
)

const (
	authTokenName       = "auth_token"
	authTokenMinimalLen = 32
)

func GetAuthTokenFilepath() string {
	return filepath.Join(filepath.Dir(os.ExpandEnv("${HOME}/.opsAgent/")), authTokenName)
}

func CreateOrFetchToken() (string, error) {
	return fetchAuthToken(true)
}

func fetchAuthToken(tokenCreationAllowed bool) (string, error) {
	authTokenFile := GetAuthTokenFilepath()

	if _, e := os.Stat(authTokenFile); os.IsNotExist(e) && tokenCreationAllowed {
		key := make([]byte, authTokenMinimalLen)
		_, e = rand.Read(key)
		if e != nil {
			return "", fmt.Errorf("can't create agent authentication token value: %s", e)
		}

		e = saveAuthToken(hex.EncodeToString(key), authTokenFile)
		if e != nil {
			return "", fmt.Errorf("error writing authentication token file on fs: %s", e)
		}
		log.Infof("Saved a new authentication token to %s", authTokenFile)
	}

	authTokenRaw, e := ioutil.ReadFile(authTokenFile)
	if e != nil {
		return "", fmt.Errorf("unable to read authentication token file: " + e.Error())
	}

	authToken := strings.TrimSpace(string(authTokenRaw))
	if len(authToken) < authTokenMinimalLen {
		return "", fmt.Errorf("invalid authentication token: must be at least %d characters in length", authTokenMinimalLen)
	}

	return authToken, nil
}

func saveAuthToken(token, tokenPath string) error {
	if err := ioutil.WriteFile(tokenPath, []byte(token), 0o600); err != nil {
		return err
	}

	perms, err := filesystem.NewPermission()
	if err != nil {
		return err
	}

	if err := perms.RestrictAccessToUser(tokenPath); err != nil {
		log.Errorf("Failed to write auth token acl %s", err)
		return err
	}

	log.Infof("Wrote auth token")
	return nil
}
