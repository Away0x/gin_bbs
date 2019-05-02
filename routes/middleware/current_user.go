package middleware

import (
	"gin_bbs/app/auth"

	"github.com/gin-gonic/gin"
)

// CurrentUserMiddleware : 从 session 中获取 user model 的 middleware
func CurrentUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth.SaveCurrentUserToContext(c)

		c.Next()
	}
}
