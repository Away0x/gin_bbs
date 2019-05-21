package permission

import (
	"fmt"
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
	if val, ok := GetUserPermissionCache(u.ID, permissionName); ok {
		return val
	}

	joinSQL := fmt.Sprintf(`INNER JOIN %s as p ON p.name = '%s'`,
		(Permission{}).TableName(), permissionName)
	joinSQL2 := fmt.Sprintf(`INNER JOIN %s as m ON
    m.model_id = %d And m.model_type = '%s' AND %s.role_id = m.role_id AND %s.permission_id = p.id`,
		(ModelHasRole{}).TableName(),
		u.ID,
		u.TableName(),
		(RoleHasPermission{}).TableName(),
		(RoleHasPermission{}).TableName())

	rhp := &RoleHasPermission{}
	err := database.DB.Joins(joinSQL).Joins(joinSQL2).First(&rhp).Error
	if err != nil || rhp == nil {
		SetUserPermissionCache(u.ID, permissionName, false)
		return false
	}

	SetUserPermissionCache(u.ID, permissionName, true)
	return true
}

// GetUserAllPermission 获取用户所有权限
func GetUserAllPermission(u *userModel.User) ([]*Permission, error) {
	var (
		modelRs = make([]*ModelHasRole, 0)
		rhpS    = make([]*RoleHasPermission, 0)
		ps      = make([]*Permission, 0)

		roleIDs = make([]uint, 0) // 存储所有角色 id
		pIDs    = make([]uint, 0) // 存储所有权限 id
	)

	// 获取用户所有角色
	if err := database.DB.Where("model_type = ? AND model_id = ?", u.TableName(), u.ID).Find(&modelRs).Error; err != nil {
		return ps, err
	}
	for _, v := range modelRs {
		roleIDs = append(roleIDs, v.RoleID)
	}
	// 获取角色的所有权限
	if err := database.DB.Where("role_id in (?)", roleIDs).Find(&rhpS).Error; err != nil {
		return ps, err
	}
	for _, v := range rhpS {
		pIDs = append(pIDs, v.PermissionID)
	}
	// 获取权限的所有信息
	if err := database.DB.Where("id in (?)", pIDs).Find(&ps).Error; err != nil {
		return ps, err
	}

	return ps, nil
}
