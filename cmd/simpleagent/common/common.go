package common

import (
	"context"
	"encoding/json"
	"net/http"
	"simpleagent/pkg/version"
)

var (
	// MainCtx is the main agent context passed to components
	MainCtx context.Context

	// MainCtxCancel cancels the main agent context
	MainCtxCancel context.CancelFunc
)

// GetVersion returns the version of the agent in a http response json
func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	av, _ := version.Agent()
	j, _ := json.Marshal(av)
	w.Write(j)
}
