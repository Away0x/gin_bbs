package login

import (
	"gin_bbs/app/auth"
	"gin_bbs/app/controllers"
	"gin_bbs/pkg/ginutils/flash"

	userRequest "gin_bbs/app/requests/user"
	"gin_bbs/pkg/ginutils/validate"

	"github.com/gin-gonic/gin"
)

func ShowLoginForm(c *gin.Context) {
	controllers.Render(c, "auth/login", gin.H{})
}

func Login(c *gin.Context) {
	// 验证参数并且获取用户
	userLoginForm := &userRequest.UserLoginForm{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}
	user, errors := userLoginForm.ValidateAndGetUser(c)

	if len(errors) != 0 || user == nil {
		validate.SaveValidateMessage(c, errors)
		controllers.RedirectRouter(c, "login")
		return
	}

	auth.Login(c, user)
	controllers.RedirectRouter(c, "root")
}

func Logout(c *gin.Context) {
	auth.Logout(c)
	flash.NewSuccessFlash(c, "您已成功退出！")
	controllers.RedirectRouter(c, "login")
}
