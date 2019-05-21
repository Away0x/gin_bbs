package cron

import (
	"github.com/robfig/cron"
)

var (
	cronActionMap = map[string]func(){}
)

func init() {
	// 记录活跃用户
	cronActionMap["0 0 0/1 * * *"] = ActiveUser
	// 用户最后登录时间
	cronActionMap["0 0 0/1 * * *"] = LastActivedUser
}

// New -
func New() *cron.Cron {
	c := cron.New()
	for spec, cmd := range cronActionMap {
		c.AddFunc(spec, cmd)
	}

	return c
}
