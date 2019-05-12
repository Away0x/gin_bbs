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

	setToCache(t)
	return nil
}

// Update -
func (t *Topic) Update() (err error) {
	if err = database.DB.Save(&t).Error; err != nil {
		log.Warnf("topic 更新失败: %v", err)
		return err
	}

	setToCache(t)
	return nil
}

// Delete -
func Delete(id int) (err error) {
	if err = database.DB.Where("id = ?", id).Delete(&Topic{}).Error; err != nil {
		log.Warnf("topic 删除失败: %v", err)
		return err
	}

	delCache(id)
	return nil
}
