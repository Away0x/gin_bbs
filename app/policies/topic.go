package policies

import (
	userModel "gin_bbs/app/models/user"

	"github.com/gin-gonic/gin"
)

// TopicPolicyUpdate : 是否有更新 topic 的权限
func TopicPolicyUpdate(c *gin.Context, currentUser *userModel.User, targetUserID int) bool {
	if currentUser.ID != uint(targetUserID) {
		Unauthorized(c)
		return false
	}

	return true
}
