package cron

import (
	"fmt"
	"gin_bbs/app/helpers"
)

// ActiveUser 记录活跃用户
func ActiveUser() {
	a := new(helpers.ActiveUser)
	fmt.Println("记录活跃用户")
	a.CalculateAndCacheActiveUsers()
}
