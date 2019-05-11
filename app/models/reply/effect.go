package reply

import (
	"gin_bbs/database"

	"github.com/lexkong/log"
)

// Create -
func (r *Reply) Create() (err error) {
	if err = database.DB.Create(&r).Error; err != nil {
		log.Warnf("reply 创建失败: %v", err)
		return err
	}

	return nil
}

// Update -
func (r *Reply) Update() (err error) {
	if err = database.DB.Save(&r).Error; err != nil {
		log.Warnf("reply 更新失败: %v", err)
		return err
	}

	return nil
}

// Delete -
func Delete(id int) (err error) {
	reply := &Reply{}
	reply.ID = uint(id)

	if err = database.DB.Delete(&reply).Error; err != nil {
		log.Warnf("reply 删除失败: %v", err)
		return err
	}

	return nil
}
