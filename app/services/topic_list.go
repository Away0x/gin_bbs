package services

import (
	categoryModel "gin_bbs/app/models/category"
	topicModel "gin_bbs/app/models/topic"
	userModel "gin_bbs/app/models/user"
	viewmodels "gin_bbs/app/viewmodels"
	"gin_bbs/database"
	"gin_bbs/pkg/ginutils/utils"
)

// TopicListService -
func TopicListService(getToipcsFunc func() ([]*topicModel.Topic, error)) (interface{}, error) {
	var (
		result = make([]interface{}, 0) // 最终结果

		topicIDs   = make([]uint, 0) // 存储所有 topic id (用于最后排序 topic)
		topicIDMap = make(map[uint]map[string]interface{})

		usersIDs = make([]uint, 0) // 存储所有 user id
		users    = make([]*userModel.User, 0)

		catIDs = make([]uint, 0) // 存储所有 category id
		cats   = make([]*categoryModel.Category, 0)
	)

	// 获取 topic
	topics, err := getToipcsFunc()
	if err != nil {
		return nil, err
	}

	for _, t := range topics {
		topicIDMap[t.ID] = viewmodels.NewTopicViewModelSerializer(t)
		topicIDs = append(topicIDs, t.ID)
		usersIDs = append(usersIDs, t.UserID)
		catIDs = append(catIDs, t.CategoryID)
	}

	// 获取 user 和 category
	if err = database.DB.Where("id in (?)", utils.UniqueUintSlice(usersIDs)).Find(&users).Error; err != nil {
		return nil, err
	}
	if err = database.DB.Where("id in (?)", utils.UniqueUintSlice(catIDs)).Find(&cats).Error; err != nil {
		return nil, err
	}

	// 整理数据
	for _, t := range topics {
		for _, u := range users {
			if t.UserID == u.ID {
				topicIDMap[t.ID]["User"] = viewmodels.NewUserViewModelSerializer(u)
			}
		}

		for _, c := range cats {
			if t.CategoryID == c.ID {
				topicIDMap[t.ID]["Category"] = c
			}
		}
	}

	for _, id := range topicIDs {
		result = append(result, topicIDMap[id])
	}

	return result, nil
}
