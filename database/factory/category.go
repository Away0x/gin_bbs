package factory

import (
	"fmt"
	"gin_bbs/app/models/category"
)

// CategoryTableSeeder -
func CategoryTableSeeder(needCleanTable bool) {
	if needCleanTable {
		dropAndCreateTable(&category.Category{})
	}

	cs := []*category.Category{
		{
			Name:        "分享",
			Description: "分享创造，分享发现",
		},
		{
			Name:        "教程",
			Description: "开发技巧、推荐扩展包等",
		},
		{
			Name:        "问答",
			Description: "请保持友善，互帮互助",
		},
		{
			Name:        "公告",
			Description: "站点公告",
		},
	}

	for _, c := range cs {
		if err := c.Create(); err != nil {
			fmt.Printf("mock category error： %v\n", err)
		}
	}
}
