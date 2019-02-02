package storage

import "github.com/karlseguin/ccache"

type GlobalMemoryCache struct {
	Mcache *ccache.Cache
}

var GlobalCache *GlobalMemoryCache

func init() {
	if GlobalCache == nil {
		GlobalCache = new(GlobalMemoryCache)
		GlobalCache.Mcache = ccache.New(ccache.Configure().MaxSize(100000).ItemsToPrune(5000))
	}
}

func GetGlobalCache() *GlobalMemoryCache {
	return GlobalCache
}
