package user

import (
	"gin_bbs/database"

	"github.com/lexkong/log"
)

// Create -
func (u *User) Create() (err error) {
	if err = database.DB.Create(&u).Error; err != nil {
		log.Warnf("用户创建失败: %v", err)
		return err
	}

	setToCache(u)
	return nil
}

// Update 更新用户
func (u *User) Update() (err error) {
	if err = database.DB.Model(&User{}).Updates(&u).Error; err != nil {
		log.Warnf("用户更新失败: %v", err)
		return err
	}

	setToCache(u)
	return nil
}

// Notification 更新用户通知
func (u *User) Notification(count int) (err error) {
	if err = database.DB.Model(&User{}).Updates(map[string]interface{}{
		"notification_count": count,
	}).Error; err != nil {
		return err
	}
	u.NotificationCount = 0

	setToCache(u)
	return nil
}

// Delete -
func Delete(id int) (err error) {
	user := &User{}
	user.BaseModel.ID = uint(id)

	// Unscoped: 永久删除而不是软删除 (由于该操作是管理员操作的，所以不使用软删除)
	if err = database.DB.Unscoped().Delete(&user).Error; err != nil {
		log.Warnf("用户删除失败: %v", err)
		return err
	}

	delCache(id)
	return nil
}
