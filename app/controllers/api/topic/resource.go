package topic

import (
	"gin_bbs/app/controllers"
	userModel "gin_bbs/app/models/user"
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

}
