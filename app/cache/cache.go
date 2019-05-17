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

// PutInt64 -
func PutInt64(key string, val int64, expiration time.Duration) {
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

// GetInt64 -
func GetInt64(key string) (int64, bool) {
	d, ok := Get(key)
	if !ok {
		return 0, false
	}
	if d, ok := d.(int64); ok {
		return d, true
	}

	return 0, false
}

// Del -
func Del(key string) {
	cacheInstance.Delete(key)
}
