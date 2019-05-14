package captcha

import (
	"gin_bbs/app/controllers"
	"gin_bbs/app/requests"

	"github.com/gin-gonic/gin"
)

// Store 图片验证码
func Store(c *gin.Context) {
	phone, ok := requests.RunPhoneValidate(c)
	if !ok {
		return
	}

	controllers.SendOKResponse(c, phone)
}
