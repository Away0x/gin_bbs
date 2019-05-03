package register

import (
	"gin_bbs/app/auth"
	"gin_bbs/app/controllers"
	userRequest "gin_bbs/app/requests/user"

	"gin_bbs/pkg/ginutils/captcha"
	"gin_bbs/pkg/ginutils/flash"

	"gin_bbs/app/helpers"

	"github.com/gin-gonic/gin"
)

// 展示注册页面
func ShowRegistrationForm(c *gin.Context) {
	captcha := captcha.New("/captcha")

	controllers.Render(c, "auth/register", gin.H{
		"captcha": captcha,
	})
}

// 注册
func Register(c *gin.Context) {
	// 验证参数和创建用户
	userCreateForm := &userRequest.UserCreateForm{
		Name:                 c.PostForm("name"),
		Email:                c.PostForm("email"),
		Password:             c.PostForm("password"),
		PasswordConfirmation: c.PostForm("password_confirmation"),
		Captcha:              c.PostForm("captcha"),
		CaptchaID:            c.PostForm("captcha_id"),
	}
	ok, user := userCreateForm.ValidateAndSave(c)

	if !ok || user == nil {
		controllers.RedirectRouter(c, "register")
		return
	}

	auth.Login(c, user)
	if err := helpers.SendVerifyEmail(user); err != nil {
		flash.NewDangerFlash(c, "邮件发送失败: "+err.Error())
	}
	controllers.RedirectRouter(c, "root")
}
