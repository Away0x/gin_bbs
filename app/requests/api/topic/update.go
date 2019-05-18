package topic

import (
	topicModel "gin_bbs/app/models/topic"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils/validate"
)

// Update 更新 topic
type Update struct {
	validate.Validate
	Title      string `json:"title"`
	Body       string `json:"body"`
	CategoryID uint   `json:"category_id"`
}

// RegisterValidators 注册验证器
func (u *Update) RegisterValidators() validate.ValidatorMap {
	return validate.ValidatorMap{
		"category_id": {
			categoryIDValidator(u.CategoryID),
		},
	}
}

// Run -
func (u *Update) Run(topic *topicModel.Topic) *errno.Errno {
	ok, _, errMap := validate.Run(u)
	if !ok {
		return errno.New(errno.ParamsError, errMap)
	}

	if u.Title != "" {
		topic.Title = u.Title
	}
	if u.Body != "" {
		topic.Body = u.Body
	}
	topic.CategoryID = u.CategoryID

	if err := topic.Update(); err != nil {
		return errno.New(errno.DatabaseError, err)
	}

	return nil
}
