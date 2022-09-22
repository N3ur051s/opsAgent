package util

import (
	"fmt"
	"net/http"
	"simpleagent/pkg/api/security"
	"strings"
)

var (
	token string
)

func CreateAndSetAuthToken() error {
	// Noop if token is already set
	if token != "" {
		return nil
	}

	// token is only set once, no need to mutex protect
	var err error
	token, err = security.CreateOrFetchToken()
	return err
}

func GetAuthToken() string {
	return token
}

// Validate validates an http request
func Validate(w http.ResponseWriter, r *http.Request) error {
	var err error
	auth := r.Header.Get("Authorization")
	if auth == "" {
		w.Header().Set("WWW-Authenticate", `Bearer realm="SimpleAgent"`)
		err = fmt.Errorf("no session token provided")
		http.Error(w, err.Error(), 401)
		return err
	}

	tok := strings.Split(auth, " ")
	if tok[0] != "Bearer" {
		w.Header().Set("WWW-Authenticate", `Bearer realm="SimpleAgent"`)
		err = fmt.Errorf("unsupported authorization scheme: %s", tok[0])
		http.Error(w, err.Error(), 401)
		return err
	}

	if len(tok) < 2 || tok[1] != GetAuthToken() {
		err = fmt.Errorf("invalid session token")
		http.Error(w, err.Error(), 403)
	}

	return err
}
