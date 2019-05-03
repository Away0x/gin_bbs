package verification

import (
	"gin_bbs/app/controllers"
	"gin_bbs/app/helpers"
	userModel "gin_bbs/app/models/user"

	"gin_bbs/pkg/ginutils/flash"

	"github.com/gin-gonic/gin"
)

// 展示发送激活邮件的页面
func Show(c *gin.Context) {
	controllers.Render(c, "auth/verify", gin.H{})
}

// 激活
func Verify(c *gin.Context) {

}

// 重新发送激活邮件
func Resend(c *gin.Context, currentUser *userModel.User) {
	if err := helpers.SendVerifyEmail(currentUser); err != nil {
		flash.NewDangerFlash(c, "邮件发送失败: "+err.Error())
	}
	controllers.RedirectRouter(c, "verification.notice")
}
