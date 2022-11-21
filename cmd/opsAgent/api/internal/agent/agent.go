package agent

type contextKey struct {
	key string
}

// ConnContextKey key to reference the http connection from the request context
var ConnContextKey = &contextKey{"http-connection"}
