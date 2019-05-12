package policies

import (
	"gin_bbs/app/controllers"

	"github.com/gin-gonic/gin"
)

// Unauthorized : 无权限时
func Unauthorized(c *gin.Context) {
	controllers.RenderUnauthorized(c)
}

// CheckPolicy 检查权限
func CheckPolicy(c *gin.Context, cond func() bool) bool {
	if cond() {
		return true
	}

	Unauthorized(c)
	return false
}
