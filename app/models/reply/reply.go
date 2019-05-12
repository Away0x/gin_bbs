package reply

import (
	"gin_bbs/app/models"
	"gin_bbs/app/models/topic"
	"gin_bbs/pkg/ginutils/utils"
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

// BeforeSave - hook
func (r *Reply) BeforeSave() error {
	r.Content = utils.XSSClean(r.Content)

	return nil
}

// AfterCreate - hook
func (r *Reply) AfterCreate() (err error) {
	// 更新 topic 的 reply count
	t, err := topic.Get(int(r.TopicID))
	if err != nil {
		return err
	}
	// 注意，在 GORM 中的保存 / 删除 操作会默认进行事务处理，所以在事物中，所有的改变都是无效的，直到它被提交为止
	// 所以该钩子中，获取的 count 是创建前的，需要自己加 1
	newCount, err := CountByTopicID(int(t.ID))
	if err != nil {
		return err
	}

	t.ReplyCount = newCount + 1
	if err := t.Update(); err != nil {
		return err
	}

	return nil
}
