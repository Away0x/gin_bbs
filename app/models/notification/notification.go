package notification

import (
	"time"
)

// Notification -
type Notification struct {
	ID             uint       `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Type           string     `gorm:"column:type;type:varchar(255);not null" sql:"index"` // notification 类型 (作用)
	NotifiableType string     `gorm:"column:notifiable_type;type:varchar(255);not null"`  // 接收者的目标类型
	NotifiableID   uint       `gorm:"column:notifiable_id;not null" sql:"index"`          // 接收者的标识 (id)
	Data           string     `gorm:"column:data;type:text;not null"`                     // 发送的数据
	ReadAt         *time.Time `gorm:"column:read_at"`                                     // 读消息的时间
	CreatedAt      time.Time  `gorm:"column:created_at"`
	UpdatedAt      time.Time  `gorm:"column:updated_at"`
}

// TableName 表名
func (Notification) TableName() string {
	return "notifications"
}
