package permission

import (
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	permissionCache = cache.New(30*time.Minute, 1*time.Hour)
)

func getCacheKey(userid uint, permissionName string) string {
	return strconv.Itoa(int(userid)) + "-" + permissionName
}

// GetUserPermissionCache -
func GetUserPermissionCache(userid uint, permissionName string) (bool, bool) {
	key := getCacheKey(userid, permissionName)
	status, ok := permissionCache.Get(key)
	if !ok {
		return false, false
	}

	s, ok := status.(bool)
	if !ok {
		return false, false
	}

	return s, true
}

// SetUserPermissionCache -
func SetUserPermissionCache(userid uint, permissionName string, status bool) {
	key := getCacheKey(userid, permissionName)
	permissionCache.Set(key, status, cache.DefaultExpiration)
}
