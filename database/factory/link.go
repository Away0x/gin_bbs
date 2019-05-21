package factory

import (
	"fmt"
	"gin_bbs/app/models/link"
	"strconv"
)

func linkFactory(i int) *link.Link {
	index := strconv.Itoa(i)
	return &link.Link{
		Title: "link title " + index,
		Link:  "https://www.baidu.com/" + index,
	}
}

// LinksTableSeeder -
func LinksTableSeeder(needCleanTable bool) {
	if needCleanTable {
		dropAndCreateTable(&link.Link{})
	}

	for i := 0; i < 6; i++ {
		link := linkFactory(i)
		if err := link.Create(); err != nil {
			fmt.Printf("mock link error： %v\n", err)
		}
	}
}
