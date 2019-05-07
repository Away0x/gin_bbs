package topic

import (
	"gin_bbs/database"

	"github.com/lexkong/log"
)

// Create -
func (t *Topic) Create() (err error) {
	if err = database.DB.Create(&t).Error; err != nil {
		log.Warnf("topic 创建失败: %v", err)
		return err
	}

	return nil
}
