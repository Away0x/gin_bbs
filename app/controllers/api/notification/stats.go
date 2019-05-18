package notification

import (
	"gin_bbs/app/controllers"
	userModel "gin_bbs/app/models/user"

	"github.com/gin-gonic/gin"
)

// Stats 通知统计
func Stats(c *gin.Context, currentUser *userModel.User, tokenString string) {
	controllers.SendOKResponse(c, map[string]int{
		"unread_count": currentUser.NotificationCount,
	})
}
