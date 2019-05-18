package reply

import (
	"gin_bbs/app/controllers"
	replyModel "gin_bbs/app/models/reply"
	topicModel "gin_bbs/app/models/topic"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/app/services"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils"

	"github.com/gin-gonic/gin"
)

// TopicReplies topic 的 reply list
func TopicReplies(c *gin.Context) {
	topicID, err := ginutils.GetIntParam(c, "topic_id")
	if err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}

	topic, err := topicModel.Get(topicID)
	if err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.DatabaseError, err))
		return
	}

	// 获取回复
	replies, _ := services.RpleyListApiService(func() ([]*replyModel.Reply, error) {
		return replyModel.TopicReplies(int(topic.ID))
	})

	controllers.SendOKResponse(c, controllers.ListData{List: replies})
}

// UserReplies user 的 reply list
func UserReplies(c *gin.Context) {
	userID, err := ginutils.GetIntParam(c, "user_id")
	if err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}

	user, err := userModel.Get(userID)
	if err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.DatabaseError, err))
		return
	}

	// 获取回复
	replies, _ := services.RpleyListApiService(func() ([]*replyModel.Reply, error) {
		return replyModel.UserReplies(int(user.ID), 0, 0)
	})

	controllers.SendOKResponse(c, controllers.ListData{List: replies})
}
