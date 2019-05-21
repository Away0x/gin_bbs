package permission

import (
	"gin_bbs/database"

	"github.com/lexkong/log"
)

// RoleHasPermission 角色拥有的权限关联表，如管理员拥有查看后台的权限都是在此表定义，一个角色能拥有多个权限
type RoleHasPermission struct {
	PermissionID uint `gorm:"column:permission_id;not null"`
	RoleID       uint `gorm:"column:role_id;not null" sql:"index"`
}

// TableName 表名
func (RoleHasPermission) TableName() string {
	return "role_has_permissions"
}

// Create -
func (r *RoleHasPermission) Create() (err error) {
	if err = database.DB.Create(&r).Error; err != nil {
		log.Warnf("RoleHasPermission 创建失败: %v", err)
		return err
	}

	return nil
}
