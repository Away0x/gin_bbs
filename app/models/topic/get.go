package topic

import (
	"gin_bbs/app/models/category"
	"gin_bbs/app/models/user"
	"gin_bbs/database"
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
func List(offset, limit int) (topics []*Topic, err error) {
	topics = make([]*Topic, 0)

	if err = database.DB.Offset(offset).Limit(limit).Order("created_at desc").Find(&topics).Error; err != nil {
		return topics, err
	}

	return topics, nil
}

// Count -
func Count() (count int, err error) {
	err = database.DB.Model(&Topic{}).Count(&count).Error
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
func GetByCategoryID(categoryID, offset, limit int) (topics []*Topic, err error) {
	topics = make([]*Topic, 0)

	if err = database.DB.Where("category_id = ?", categoryID).Offset(offset).Limit(limit).Order("created_at desc").Find(&topics).Error; err != nil {
		return topics, err
	}

	return topics, nil
}

// Category 获取 topic 的 category
func Category(topicID int) (*category.Category, error) {
	t, err := Get(topicID)
	if err != nil {
		return nil, err
	}

	cat, err := category.Get(int(t.CategoryID))
	if err != nil {
		return nil, err
	}

	return cat, err
}

// User 获取 topic 的 user
func User(topicID int) (*user.User, error) {
	t, err := Get(topicID)
	if err != nil {
		return nil, err
	}

	u, err := user.Get(int(t.UserID))
	if err != nil {
		return nil, err
	}

	return u, err
}
