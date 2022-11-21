package util

import (
	"opsAgent/pkg/api/security"
)

var (
	token string
)

func CreateAndSetAuthToken() error {
	if token != "" {
		return nil
	}

	var err error
	token, err = security.CreateOrFetchToken()

	if err != nil {
		return err
	}
	return nil
}

func GetAuthToken() string {
	return token
}
