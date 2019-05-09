package topic

import (
	"gin_bbs/app/models/category"
	"gin_bbs/app/models/user"
	"gin_bbs/database"
)

const (
	orderDefault = "default"
	orderRecent  = "recent"
)

// Get -
func Get(id int) (*Topic, error) {
	t := &Topic{}
	if err := database.DB.First(&t, id).Error; err != nil {
		return t, nil
	}

	return t, nil
}

// List -
func List(offset, limit int, order string) (topics []*Topic, err error) {
	topics = make([]*Topic, 0)

	if order == orderRecent {
		order = "created_at"
	} else {
		order = "updated_at"
	}

	if err = database.DB.Offset(offset).Limit(limit).Order(order + " desc").Find(&topics).Error; err != nil {
		return topics, err
	}

	return topics, nil
}

// Count -
func Count() (count int, err error) {
	err = database.DB.Model(&Topic{}).Count(&count).Error
	return
}

// CountByCategoryID -
func CountByCategoryID(categoryID int) (count int, err error) {
	err = database.DB.Model(&Topic{}).Where("category_id = ?", categoryID).Count(&count).Error
	return
}

// CountByUserID -
func CountByUserID(userID int) (count int, err error) {
	err = database.DB.Model(&Topic{}).Where("user_id = ?", userID).Count(&count).Error
	return
}

// All -
func All() (topics []*Topic, err error) {
	topics = make([]*Topic, 0)

	if err = database.DB.Order("created_at desc").Find(&topics).Error; err != nil {
		return topics, err
	}

	return topics, nil
}

// GetByCategoryID 根据 category_id 获取 topics
func GetByCategoryID(categoryID, offset, limit int, order string) (topics []*Topic, err error) {
	topics = make([]*Topic, 0)

	if order == orderRecent {
		order = "created_at"
	} else {
		order = "updated_at"
	}

	if err = database.DB.Where("category_id = ?", categoryID).Offset(offset).Limit(limit).Order(order + " desc").Find(&topics).Error; err != nil {
		return topics, err
	}

	return topics, nil
}

// GetByUserID -
func GetByUserID(userID, offset, limit int) (topics []*Topic, err error) {
	topics = make([]*Topic, 0)

	if err = database.DB.Where("user_id = ?", userID).Offset(offset).Limit(limit).Order("created_at desc").Find(&topics).Error; err != nil {
		return topics, err
	}

	return topics, nil
}

// TopicAndCategory 获取 topic 的 category
func TopicAndCategory(topicID int) (*Topic,*category.Category, error) {
	t, err := Get(topicID)
	if err != nil {
		return nil, nil, err
	}

	c, err := category.Get(int(t.CategoryID))
	if err != nil {
		return nil, nil, err
	}

	return t, c, err
}

// TopicAndUser 获取 topic 的 user
func TopicAndUser(topicID int) (*Topic, *user.User, error) {
	t, err := Get(topicID)
	if err != nil {
		return nil, nil, err
	}

	u, err := user.Get(int(t.UserID))
	if err != nil {
		return nil, nil, err
	}

	return t, u, err
}
