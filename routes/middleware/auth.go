package middleware

import (
	"gin_bbs/app/auth"
	"gin_bbs/app/controllers"
	"gin_bbs/app/helpers"

	"github.com/gin-gonic/gin"
)

// Auth 用户已登录才能访问
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, err := auth.GetCurrentUserFromContext(c)
		if currentUser == nil || err != nil {
			controllers.RedirectRouter(c, "login")
			c.Abort()
			return
		}

		helpers.RecordLastActivedAt(currentUser)
		c.Next()
	}
}
