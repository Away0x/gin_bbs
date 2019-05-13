package reply

import (
	"errors"
	notification "gin_bbs/app/models/notification"
	topicModel "gin_bbs/app/models/topic"
	userModel "gin_bbs/app/models/user"
	"strconv"
	// "gin_bbs/app/helpers"
)

// TopicRepliedNotify 发送一条消息给 topic 作者 (告知其有新回复了)
func TopicRepliedNotify(reply *Reply, currentUser *userModel.User) error {
	data, topic, err := getTopicReplied(reply, currentUser)
	if err != nil {
		return err
	}

	if err = notification.Notify("TopicReplied", userModel.TableName, topic.UserID, data); err != nil {
		return err
	}

	// user notification_count ++ (topic 作者的消息 count +1)
	topicAuthor, err := userModel.Get(int(topic.UserID))
	if err != nil {
		return err
	}
	topicAuthor.NotificationCount = topicAuthor.NotificationCount + 1
	err = topicAuthor.Update()

	// 发送通知邮件 (暂时关闭)
	// go func() {
	// 	helpers.SendMail([]string{topicAuthor.Email},
	// 		"Topic Replied",
	// 		"mail/notifications/topic_replied.html",
	// 		map[string]interface{}{	"URL": data["topic_link"]})
	// }()

	return err
}

// ----------------------- private
func getTopicReplied(reply *Reply, currentUser *userModel.User) (map[string]interface{}, *topicModel.Topic, error) {
	topic, err := topicModel.Get(int(reply.TopicID))
	if err != nil {
		return nil, nil, err
	}
	// 要通知的人是当前用户 (帖子主人正是当前发布 reply 的用户)
	if topic.UserID == currentUser.ID {
		return nil, nil, errors.New("帖子主人正是当前发布评论的用户")
	}

	user, err := userModel.Get(int(reply.UserID))
	if err != nil {
		return nil, nil, err
	}

	data := map[string]interface{}{
		"reply_id":      reply.ID,
		"reply_content": reply.Content,
		"user_id":       reply.UserID,
		"user_name":     user.Name,
		"user_avatar":   user.Avatar,
		"topic_link":    topic.Link() + "#reply" + strconv.Itoa(int(reply.ID)),
		"topic_id":      topic.ID,
		"topic_title":   topic.Title,
	}

	return data, topic, nil
}
