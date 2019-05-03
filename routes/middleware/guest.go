package middleware

import (
	"gin_bbs/app/auth"
	"gin_bbs/app/controllers"
	"gin_bbs/pkg/ginutils/flash"

	"github.com/gin-gonic/gin"
)

// Guest 用户未登录才能访问
func Guest() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, err := auth.GetCurrentUserFromContext(c)
		if currentUser != nil && err == nil {
			flash.NewInfoFlash(c, "您已登录，无需再次操作。")
			controllers.RedirectRouter(c, "root")
			c.Abort()
			return
		}

		c.Next()
	}
}
