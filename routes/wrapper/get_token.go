package wrapper

import (
	"gin_bbs/app/auth/token"
	"gin_bbs/app/controllers"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/pkg/errno"

	"github.com/gin-gonic/gin"
)

// GetToken 获取 token
func GetToken(handler func(*gin.Context, *userModel.User, string)) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, user, ok := token.GetTokenUserFromContext(c)
		if !ok {
			controllers.SendErrorResponse(c, errno.TokenError)
			return
		}

		handler(c, user, tokenStr)
	}
}
