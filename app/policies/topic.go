package policies

import (
	userModel "gin_bbs/app/models/user"

	"github.com/gin-gonic/gin"
)

// TopicPolicyOwner : 是否有更新、删除 topic 的权限
func TopicPolicyOwner(c *gin.Context, currentUser *userModel.User, targetUserID int) bool {
	if CheckTopicPolicyOwner(currentUser, targetUserID) {
		return true
	}

	Unauthorized(c)
	return false
}

// CheckTopicPolicyOwner -
func CheckTopicPolicyOwner(currentUser *userModel.User, targetUserID int) bool {
	if currentUser == nil {
		return false
	}
	if before(currentUser) {
		return true
	}

	if currentUser.ID != uint(targetUserID) {
		return false
	}

	return true
}
