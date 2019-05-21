package policies

import (
	userModel "gin_bbs/app/models/user"

	"github.com/gin-gonic/gin"
)

// UserPolicyUpdate : 是否有更新目标 user 的权限
func UserPolicyUpdate(c *gin.Context, currentUser *userModel.User, targetUserID int) bool {
  if currentUser == nil {
    return false
  }

  if before(currentUser) {
		return true
	}

	if currentUser.ID != uint(targetUserID) {
		Unauthorized(c)
		return false
	}

	return true
}
