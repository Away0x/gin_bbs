package reply

import (
	"gin_bbs/app/models/topic"
	"gin_bbs/app/models/user"
	"gin_bbs/database"
)

// Get -
func Get(id int) (*Reply, error) {
	r := &Reply{}
	if err := database.DB.First(&r, id).Error; err != nil {
		return r, nil
	}

	return r, nil
}

// Topic 获取回复的 topic
func (r *Reply) Topic() (t *topic.Topic, err error) {
	t, err = topic.Get(int(r.TopicID))
	if err != nil {
		return nil, err
	}

	return t, err
}

// User 获取回复的 user
func (r *Reply) User() (u *user.User, err error) {
	u, err = user.Get(int(r.UserID))
	if err != nil {
		return nil, err
	}

	return u, err
}

// TopicReplies 获取 topic 下的所有回复
func TopicReplies(topicID int) ([]*Reply, error) {
	rs := make([]*Reply, 0)
	t, err := topic.Get(topicID)
	if err != nil {
		return rs, err
	}

	if err = database.DB.Where("topic_id = ?", t.ID).Order("id desc").Find(&rs).Error; err != nil {
		return rs, err
	}

	return rs, nil
}

// UserReplies 获取用户的所有回复
func UserReplies(userID, offset, limit int) ([]*Reply, error) {
	rs := make([]*Reply, 0)
	u, err := user.Get(userID)
	if err != nil {
		return rs, err
	}

	if err = database.DB.Where("user_id = ?", u.ID).Offset(offset).Limit(limit).Order("id desc").Find(&rs).Error; err != nil {
		return rs, err
	}

	return rs, nil
}

// CountByUserID -
func CountByUserID(userID int) (count int, err error) {
	err = database.DB.Model(&Reply{}).Where("user_id = ?", userID).Count(&count).Error
	return
}
