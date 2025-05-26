package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Cache *cache.Cache

func init() {
	Cache = cache.New(cache.NoExpiration, time.Minute)
}

func Set(k string, x interface{}, d time.Duration) {
	Cache.Set(k, x, d)
}

func Get(k string) (interface{}, bool) {
	x, ok := Cache.Get(k)
	if !ok {

		return nil, false
	}

	return x, true
}
