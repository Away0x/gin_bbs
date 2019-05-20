package permission

import (
	"gin_bbs/app/models"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/database"

	"github.com/lexkong/log"
)

const (
	// RoleNameFounder 站长
	RoleNameFounder = "Founder"
	// RoleNameMaintainer 管理员
	RoleNameMaintainer = "Maintainer"
)

// Role 角色的模型表
type Role struct {
	models.BaseModel
	Name      string `gorm:"column:name;type:varchar(255);not null"`
	GuardName string `gorm:"column:guard_name;type:varchar(255);not null"`
}

// TableName 表名
func (Role) TableName() string {
	return "roles"
}

// Create 创建权限
func (r *Role) Create() (err error) {
	if err = database.DB.Create(&r).Error; err != nil {
		log.Warnf("Role 创建失败: %v", err)
		return err
	}

	return nil
}

// GetRoleByName -
func GetRoleByName(roleName string) (*Role, error) {
	r := &Role{}
	if err := database.DB.Where("name = ?", roleName).First(&r).Error; err != nil {
		return nil, err
	}

	return r, nil
}

// GivePermissionTo 赋予角色权限
func (r *Role) GivePermissionTo(permissionName string) (err error) {
	p, err := GetPermissionByName(permissionName)
	if err != nil {
		return err
	}

	rhp := &RoleHasPermission{
		PermissionID: p.ID,
		RoleID:       r.ID,
	}
	if err = rhp.Create(); err != nil {
		return err
	}

	return nil
}

// AssignRole 赋予用户角色
func (r *Role) AssignRole(u *userModel.User) (err error) {
	mhr := &ModelHasRole{
		RoleID:    r.ID,
		ModelType: u.TableName(),
		ModelID:   u.ID,
	}
	if err := mhr.Create(); err != nil {
		return err
	}

	return nil
}

// UserHasRole 用户是否是某个角色
func UserHasRole(u *userModel.User, roleName string) bool {
	r, err := GetRoleByName(roleName)
	if err != nil {
		return false
	}

	result := &ModelHasRole{}
	d := database.DB.Where("role_id = ? AND model_type = ? AND model_id = ?", r.ID, u.TableName(), u.ID).First(&result)
	if d.Error != nil && result != nil {
		return false
	}

	return true
}

// UserHasAnyRole 用户是否拥有至少一个角色
func UserHasAnyRole(u *userModel.User) bool {
	var count int
	d := database.DB.Model(&ModelHasRole{}).Where("model_type = ? AND model_id = ?", u.TableName(), u.ID).Count(&count)
	if d.Error != nil && count <= 0 {
		return false
	}

	return true
}

// UserHasAllRole 用户是否拥有所有角色
func UserHasAllRole(u *userModel.User) bool {
	var allRoleCount int
	d := database.DB.Model(&Role{}).Count(&allRoleCount)
	if d.Error != nil {
		return false
	}

	var count int
	d = database.DB.Model(&ModelHasRole{}).Where("model_type = ? AND model_id = ?", u.TableName(), u.ID).Count(&count)
	if d.Error != nil && count <= 0 {
		return false
	}

	if count < allRoleCount {
		return false
	}

	return true
}
