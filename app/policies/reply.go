package policies

import (
	replyModel "gin_bbs/app/models/reply"
	topicModel "gin_bbs/app/models/topic"
	userModel "gin_bbs/app/models/user"

	"github.com/gin-gonic/gin"
)

// ReplyPolicy : 是否有更新、删除 reply 的权限
func ReplyPolicy(c *gin.Context, currentUser *userModel.User, reply *replyModel.Reply, topic *topicModel.Topic) bool {
	if CheckReplyPolicy(currentUser, reply, topic) {
		return true
	}

	Unauthorized(c)
	return false
}

// CheckReplyPolicy -
func CheckReplyPolicy(currentUser *userModel.User, reply *replyModel.Reply, topic *topicModel.Topic) bool {
	if currentUser == nil {
		return false
	}
	if before(currentUser) {
		return true
	}

	if reply.UserID == currentUser.ID || topic.UserID == currentUser.ID {
		return true
	}

	return false
}
