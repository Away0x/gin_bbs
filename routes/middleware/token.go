package middleware

import (
	"gin_bbs/app/auth/token"
	"gin_bbs/app/controllers"

	"github.com/gin-gonic/gin"
)

// TokenAuth token 验证
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := token.GetTokenFromRequest(c)
		if err != nil || tokenStr == "" {
			controllers.SendErrorResponse(c, err)
			c.Abort()
			return
		}

		_, err = token.ParseAndGetUser(c, tokenStr)
		if err != nil {
			controllers.SendErrorResponse(c, err)
			c.Abort()
			return
		}

		c.Next()
	}
}
