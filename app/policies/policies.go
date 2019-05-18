package policies

import (
	"gin_bbs/app/controllers"
	"gin_bbs/pkg/constants"
	"gin_bbs/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Unauthorized : 无权限时
func Unauthorized(c *gin.Context) {
	if constants.IsApiRequest(c) {
		controllers.SendErrorResponse(c, errno.AuthError)
		return
	}

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
