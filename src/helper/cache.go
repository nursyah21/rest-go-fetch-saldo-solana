package helper

import (
	"sync"
	"time"
)

type walletCacheEntry struct {
	Value      int
	Expiration int64
}

type apiCacheEntry struct {
	Valid      bool
	Expiration int64
}

var (
	walletCache sync.Map
	apiCache    = make(map[string]apiCacheEntry)
)

func GetCacheWallet(wallet string) (int, bool) {
	val, ok := walletCache.Load(wallet)
	if !ok {
		return 0, false
	}

	entry := val.(walletCacheEntry)
	if time.Now().Unix() > entry.Expiration {
		return 0, false
	}

	return entry.Value, true
}

func SetCacheWallet(wallet string, balance int) {
	entry := walletCacheEntry{
		Value:      balance,
		Expiration: time.Now().Add(10 * time.Second).Unix(),
	}
	walletCache.Store(wallet, entry)
}

func GetAPIKeyCache(apiKey string) bool {
	entry, found := apiCache[apiKey]
	return found && time.Now().Unix() < entry.Expiration && entry.Valid
}

func SetAPIKeyCache(apiKey string, valid bool) {
	apiCache[apiKey] = apiCacheEntry{
		Valid:      valid,
		Expiration: time.Now().Add(10 * time.Second).Unix(),
	}
}
