package topic

import (
	"gin_bbs/app/controllers"
	topicModel "gin_bbs/app/models/topic"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/app/policies"
	request "gin_bbs/app/requests/api/topic"
	"gin_bbs/app/viewmodels"
	"gin_bbs/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Store 发布 topic
func Store(c *gin.Context, currentUser *userModel.User, tokenString string) {
	var req request.Store
	if err := c.ShouldBind(&req); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}

	topic, err := req.Run(currentUser.ID)
	if err != nil {
		controllers.SendErrorResponse(c, err)
		return
	}

	controllers.SendOKResponse(c, viewmodels.TopicApi(topic))
}

// Update 修改 topic
func Update(c *gin.Context, currentUser *userModel.User, tokenString string) {
	topic, _, ok := getTopic(c, currentUser)
	if !ok {
		return
	}

	var req request.Update
	if err := c.ShouldBind(&req); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}

	if err := req.Run(topic); err != nil {
		controllers.SendErrorResponse(c, err)
		return
	}

	controllers.SendOKResponse(c, viewmodels.TopicApi(topic))
}

// Destroy 删除 topic
func Destroy(c *gin.Context, currentUser *userModel.User, tokenString string) {
	topic, id, ok := getTopic(c, currentUser)
	if !ok {
		return
	}

	// 权限
	if ok := policies.TopicPolicyOwner(c, currentUser, int(topic.UserID)); !ok {
		return
	}

	if err := topicModel.Delete(id); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.DatabaseError, err))
		return
	}

	controllers.SendOKResponse(c, map[string]interface{}{
		"id": id,
	})
}
