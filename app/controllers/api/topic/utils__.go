package topic

import (
	"gin_bbs/app/controllers"
	topicModel "gin_bbs/app/models/topic"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/app/policies"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils"

	"github.com/gin-gonic/gin"
)

// 获取要编辑的 topic
func getTopic(c *gin.Context, currentUser *userModel.User) (*topicModel.Topic, int, bool) {
	id, err := ginutils.GetIntParam(c, "id")
	if err != nil {
		controllers.SendErrorResponse(c, errno.Base(errno.ParamsError, "id 不存在"))
		return nil, id, false
	}

	topic, err := topicModel.Get(id)
	if err != nil {
		controllers.SendErrorResponse(c, errno.ResourceNotFoundError)
		return nil, id, false
	}

	// 权限
	if currentUser != nil {
		if ok := policies.TopicPolicyOwner(c, currentUser, int(topic.UserID)); !ok {
			return nil, id, false
		}
	}

	return topic, id, true
}
