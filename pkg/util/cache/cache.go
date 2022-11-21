package cache

import (
	"path"
	"time"

	cache "github.com/patrickmn/go-cache"
)

const (
	defaultExpire    = 5 * time.Minute
	defaultPurge     = 30 * time.Second
	AgentCachePrefix = "agent"
	NoExpiration     = cache.NoExpiration
)

var Cache = cache.New(defaultExpire, defaultPurge)

func BuildAgentKey(keys ...string) string {
	keys = append([]string{AgentCachePrefix}, keys...)
	return path.Join(keys...)
}
