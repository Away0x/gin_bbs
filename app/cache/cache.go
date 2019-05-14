package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	cacheInstance = cache.New(30*time.Minute, 1*time.Hour)
)

// Put -
func Put(key string, val interface{}, expiration time.Duration) {
	cacheInstance.Set(key, val, expiration)
}

// PutStringMap -
func PutStringMap(key string, val map[string]string, expiration time.Duration) {
	Put(key, val, expiration)
}

// Get -
func Get(key string) (interface{}, bool) {
	data, ok := cacheInstance.Get(key)
	if !ok {
		return nil, false
	}

	return data, true
}

// GetStringMap -
func GetStringMap(key string) (map[string]string, bool) {
	d, ok := Get(key)
	if !ok {
		return nil, false
	}
	if d, ok := d.(map[string]string); ok {
		return d, true
	}

	return nil, false
}

// Del -
func Del(key string) {
	cacheInstance.Delete(key)
}
