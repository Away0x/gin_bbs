package image

import (
	"time"
)

var (
	// Types image type
	Types = []string{"avatar", "topic"}
)

// Image -
type Image struct {
	ID        uint      `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Type      string    `gorm:"column:type;type:varchar(255);not null" sql:"index"` // 类型 (avatar,topic)
	UserID    uint      `gorm:"column:user_id;not null" sql:"index"`                // 用户 id
	Path      string    `gorm:"column:path;type:varchar(255);not null"`             // 图片路径
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName 表名
func (Image) TableName() string {
	return "images"
}
