package storage

import "github.com/karlseguin/ccache"

// GlobalMemoryCache is the model that holds the memory cache
type GlobalMemoryCache struct {
	Mcache *ccache.Cache
}

// GlobalCache is a single-instance accessor for the memory cache
var GlobalCache *GlobalMemoryCache

func init() {
	if GlobalCache == nil {
		GlobalCache = new(GlobalMemoryCache)
		GlobalCache.Mcache = ccache.New(ccache.Configure().MaxSize(100000).ItemsToPrune(5000))
	}
}

// GetGlobalCache is a helper function to get our cache storage
func GetGlobalCache() *GlobalMemoryCache {
	return GlobalCache
}
