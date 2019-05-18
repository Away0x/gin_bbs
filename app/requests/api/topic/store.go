package topic

import (
	topicModel "gin_bbs/app/models/topic"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils/validate"
)

// Store 发布 topic
type Store struct {
	validate.Validate
	Title      string `json:"title"`
	Body       string `json:"body"`
	CategoryID uint   `json:"category_id"`
}

// RegisterValidators 注册验证器
func (s *Store) RegisterValidators() validate.ValidatorMap {
	return validate.ValidatorMap{
		"title": {
			validate.RequiredValidator(s.Title),
		},
		"body": {
			validate.RequiredValidator(s.Body),
		},
		"category_id": {
			categoryIDValidator(s.CategoryID),
		},
	}
}

// RegisterMessages 注册错误信息
func (*Store) RegisterMessages() validate.MessagesMap {
	return validate.MessagesMap{
		"title": {
			"标题不可为空",
		},
		"body": {
			"话题内容不可为空",
		},
	}
}

// Run -
func (s *Store) Run(userid uint) (*topicModel.Topic, *errno.Errno) {
	ok, _, errMap := validate.Run(s)
	if !ok {
		return nil, errno.New(errno.ParamsError, errMap)
	}

	t := &topicModel.Topic{
		Title:      s.Title,
		Body:       s.Body,
		UserID:     userid,
		CategoryID: s.CategoryID,
	}
	if err := t.Create(); err != nil {
		return nil, errno.New(errno.DatabaseError, err)
	}

	return t, nil
}
