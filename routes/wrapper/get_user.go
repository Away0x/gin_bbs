package wrapper

import (
	"gin_bbs/app/auth"
	"gin_bbs/app/controllers"
	userModel "gin_bbs/app/models/user"

	"github.com/gin-gonic/gin"
)

// GetUser 获取用户
func GetUser(handler func(*gin.Context, *userModel.User)) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, err := auth.GetCurrentUserFromContext(c)
		if currentUser == nil || err != nil {
			controllers.RedirectRouter(c, "login")
			return
		}

		handler(c, currentUser)
	}
}
