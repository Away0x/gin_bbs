package link

import (
	"gin_bbs/app/models"
	"gin_bbs/database"
	"time"

	"github.com/lexkong/log"
	"github.com/patrickmn/go-cache"
)

var (
	linkCache = cache.New(30*time.Minute, 1*time.Hour)
)

// Link 资源推荐链接
type Link struct {
	models.BaseModel
	Title string `gorm:"column:title;type:varchar(255);not null" sql:"index"` // 资源的描述
	Link  string `gorm:"column:link;type:varchar(255);not null" sql:"index"`  // 资源的链接
}

// TableName 表名
func (Link) TableName() string {
	return "links"
}

// Create -
func (l *Link) Create() (err error) {
	if err = database.DB.Create(&l).Error; err != nil {
		log.Warnf("link 创建失败: %v", err)
		return err
	}

	return nil
}

// All -
func All() ([]*Link, error) {
	if cachedData, ok := linkCache.Get("link"); ok {
		list := cachedData.([]*Link)
		return list, nil
	}

	links := make([]*Link, 0)
	if err := database.DB.Order("id").Find(&links).Error; err != nil {
		return links, err
	}

	linkCache.Set("link", links, cache.DefaultExpiration)
	return links, nil
}
