package permission

import (
	"gin_bbs/app/models"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/database"

	"github.com/lexkong/log"
)

const (
	// GuardNameWeb 用户组
	GuardNameWeb = "web"
)

const (
	// PermissionNameManageContents 管理站点内容
	PermissionNameManageContents = "manage_contents"
	// PermissionNameManageUsers 管理用户
	PermissionNameManageUsers = "manage_users"
	// PermissionNameEditSettings 管理站点设置
	PermissionNameEditSettings = "edit_settings"
)

// Permission 权限的模型表
type Permission struct {
	models.BaseModel
	Name      string `gorm:"column:name;type:varchar(255);not null"`
	GuardName string `gorm:"column:guard_name;type:varchar(255);not null"`
}

// TableName 表名
func (Permission) TableName() string {
	return "permissions"
}

// Create 创建权限
func (p *Permission) Create() (err error) {
	if err = database.DB.Create(&p).Error; err != nil {
		log.Warnf("Permission 创建失败: %v", err)
		return err
	}

	return nil
}

// GetPermissionByName -
func GetPermissionByName(permissionName string) (*Permission, error) {
	p := &Permission{}
	if err := database.DB.Where("name = ?", permissionName).First(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

// AssignPermission 赋予用户权限
func (p *Permission) AssignPermission(u *userModel.User) (err error) {
	mhp := &ModelHasPermission{
		PermissionID: p.ID,
		ModelType:    u.TableName(),
		ModelID:      u.ID,
	}
	if err := mhp.Create(); err != nil {
		return err
	}

	return nil
}

// UserHasPermission 用户是否拥有某权限
func UserHasPermission(u *userModel.User, permissionName string) bool {
	p, err := GetPermissionByName(permissionName)
	if err != nil {
		return false
	}

	mhp := &ModelHasPermission{}
	err = database.DB.Where("permission_id = ? AND model_type = ? AND model_id = ?", p.ID, u.TableName(), u.ID).First(&mhp).Error
	if err != nil || mhp == nil {
		return false
	}

	return true
}

// GetUserAllPermission 获取用户所有权限
func GetUserAllPermission(u *userModel.User) ([]*Permission, error) {
	var (
		modelPs = make([]*ModelHasPermission, 0)
		ps      = make([]*Permission, 0)
		ids     = make([]uint, 0)
	)

	if err := database.DB.Where("model_type = ? AND model_id = ?", u.TableName(), u.ID).First(&modelPs).Error; err != nil {
		return ps, err
	}

	for _, v := range modelPs {
		ids = append(ids, v.PermissionID)
	}

	if err := database.DB.Where("id in (?)", ids).First(&ps).Error; err != nil {
		return ps, err
	}

	return ps, nil
}
