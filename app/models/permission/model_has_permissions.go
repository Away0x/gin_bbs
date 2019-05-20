package permission

import (
	"gin_bbs/database"

	"github.com/lexkong/log"
)

// ModelHasPermission 模型与权限关联表，一个模型能拥有多个权限
// 作用: 用户跳过角色，直接拥有权限
type ModelHasPermission struct {
	PermissionID uint   `gorm:"column:permission_id;not null"`
	ModelType    string `gorm:"column:model_type;type:varchar(255);not null" sql:"index"`
	ModelID      uint   `gorm:"column:model_id;not null" sql:"index"`
}

// TableName 表名
func (ModelHasPermission) TableName() string {
	return "model_has_permissions"
}

// Create -
func (m *ModelHasPermission) Create() (err error) {
	if err = database.DB.Create(&m).Error; err != nil {
		log.Warnf("ModelHasPermission 创建失败: %v", err)
		return err
	}

	return nil
}
