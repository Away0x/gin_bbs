package reply

import (
	"fmt"
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
	// 注意: 由于这里要触发 delete callback hook，所以先获取到再删除，否则 hook 中得不到 TopicID
	// 所以这里没用 database.DB.Where("id = ?", id).Delete(&Reply{})
	reply, err := Get(id)
	if err != nil {
		return err
	}

	if err = database.DB.Delete(&reply).Error; err != nil {
		log.Warnf("reply 删除失败: %v", err)
		return err
	}

	return nil
}

// DeleteTopicReplies 删除 topic 下的所有 reply (注意: 要不触发 callback hook)
func DeleteTopicReplies(topicID int) error {
	return database.DB.Exec(fmt.Sprintf("delete from %s where topic_id = %d", tableName, topicID)).Error
}

// DeleteUserReplies 删除 user 的所有 reply (注意: 要不触发 callback hook)
func DeleteUserReplies(userID int) error {
	return database.DB.Exec(fmt.Sprintf("delete from %s where user_id = %d", tableName, userID)).Error
}
