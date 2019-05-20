package permission

import (
	"gin_bbs/database"

	"github.com/lexkong/log"
)

// ModelHasRole 模型与角色的关联表，用户拥有什么角色在此表中定义，一个用户能拥有多个角色
type ModelHasRole struct {
	RoleID    uint   `gorm:"column:role_id;not null"`
	ModelType string `gorm:"column:model_type;type:varchar(255);not null" sql:"index"`
	ModelID   uint   `gorm:"column:model_id;not null" sql:"index"`
}

// TableName 表名
func (ModelHasRole) TableName() string {
	return "model_has_roles"
}

// Create -
func (m *ModelHasRole) Create() (err error) {
	if err = database.DB.Create(&m).Error; err != nil {
		log.Warnf("ModelHasRole 创建失败: %v", err)
		return err
	}

	return nil
}
