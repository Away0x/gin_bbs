package helpers

import (
	"fmt"
	userModel "gin_bbs/app/models/user"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	lacache = cache.New(1*time.Hour, 10*time.Hour)
)

// RecordLastActivedAt 记录用户最后登录时间
func RecordLastActivedAt(u *userModel.User) {
	key := strconv.Itoa(int(u.ID))
	val := time.Now().Unix()

	lacache.Set(key, val, 1*time.Hour)
}

// SyncUserActivedAtToDatabase 同步用户最后登录时间到数据库
func SyncUserActivedAtToDatabase() {
	caches := lacache.Items()
	fmt.Println(caches)
	for k, v := range caches {
		id, err := strconv.Atoi(k)
		if err != nil {
			continue
		}
		user, err := userModel.Get(id)
		if err != nil {
			continue
		}

		if ttint, ok := v.Object.(int64); ok {
			ttime := time.Unix(ttint, 0)
			user.LastActivedAt = &ttime
			user.Update()
		}
	}
}

// GetUserActivedLastActivedAt 获取用户最后登录时间
func GetUserActivedLastActivedAt(u *userModel.User) *time.Time {
	key := strconv.Itoa(int(u.ID))
	t, ok := lacache.Get(key)
	if ok {
		if tt, ok := t.(int64); ok {
			ttime := time.Unix(tt, 0)
			return &ttime
		}
	}

	if u.LastActivedAt != nil {
		return u.LastActivedAt
	}

	return nil
}
