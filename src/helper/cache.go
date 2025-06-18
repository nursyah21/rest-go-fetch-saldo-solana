package helper

import (
	"sync"
	"time"
)

type cacheEntry struct {
	Value      int
	Expiration int64
}

var (
	walletCache sync.Map
)

func InitCache() {
	walletCache.Range(func(key, value interface{}) bool {
		walletCache.Delete(key)
		return true
	})
}

func GetFromCache(wallet string) (int, bool) {
	val, ok := walletCache.Load(wallet)
	if !ok {
		return 0, false
	}

	entry := val.(cacheEntry)
	if time.Now().Unix() > entry.Expiration {
		return 0, false
	}

	return entry.Value, true
}

func SetToCache(wallet string, balance int) {
	entry := cacheEntry{
		Value:      balance,
		Expiration: time.Now().Add(10 * time.Second).Unix(),
	}
	walletCache.Store(wallet, entry)
}
