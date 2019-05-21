package cron

import (
	"fmt"
	"gin_bbs/app/helpers"
)

// LastActivedUser 同步记录在缓存中用户最后登录时间到数据库中
func LastActivedUser() {
	helpers.SyncUserActivedAtToDatabase()
	fmt.Println("sync LastActivedUser")
}
