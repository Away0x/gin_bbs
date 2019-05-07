package category

import (
	"gin_bbs/database"

	"github.com/lexkong/log"
)

// Create -
func (c *Category) Create() (err error) {
	if err = database.DB.Create(&c).Error; err != nil {
		log.Warnf("category 创建失败: %v", err)
		return err
	}

	return nil
}
