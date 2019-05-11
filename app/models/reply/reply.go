package reply

import (
	"gin_bbs/app/models"
)

// Reply 回复
type Reply struct {
	models.BaseModel
	TopicID uint   `gorm:"column:topic_id;not null" sql:"index"` // topic ID
	UserID  uint   `gorm:"column:user_id;not null" sql:"index"`  // 用户 ID
	Content string `gorm:"column:content;type:text;not null"`    // 回复内容
}

// TableName 表名
func (Reply) TableName() string {
	return "replies"
}
