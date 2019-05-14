package user

import (
	"gin_bbs/app/controllers"
	request "gin_bbs/app/requests/api/user"
	"gin_bbs/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Store 用户注册
func Store(c *gin.Context) {
	var req request.User
	if err := c.BindJSON(&req); err != nil {
		controllers.SendErrorResponse(c, errno.ParamsError)
		return
	}

	if err := req.ValidateAndCreateUser(); err != nil {
		controllers.SendErrorResponse(c, err)
		return
	}

	controllers.SendOKResponse(c, nil)
}
