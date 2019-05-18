package password

import (
	"gin_bbs/app/auth"
	"gin_bbs/app/controllers"
	"gin_bbs/app/helpers"
	passwordResetModel "gin_bbs/app/models/password_reset"
	passowordRequest "gin_bbs/app/requests/password"
	"gin_bbs/pkg/ginutils/flash"
	"gin_bbs/pkg/ginutils/validate"

	"github.com/gin-gonic/gin"
)

// ShowLinkRequestForm 展示发送密码重置链接的页面
func ShowLinkRequestForm(c *gin.Context) {
	controllers.Render(c, "auth/password/email", gin.H{})
}

// SendResetLinkEmail 发送密码重置链接
func SendResetLinkEmail(c *gin.Context) {
	email := c.PostForm("email")
	ok, errArr, errMap := validate.RunSingle("email",
		[]validate.ValidatorFunc{
			validate.RequiredValidator(email),
			validate.MaxLengthValidator(email, 255),
			validate.EmailValidator(email),
			emailExistValidator(email),
		},
		[]string{"邮箱不能为空", "邮箱长度不能大于 255 个字符", "邮箱格式错误"})

	if !ok {
		validate.SaveValidateMessage(c, errArr, errMap)
		controllers.RedirectRouter(c, "password.request")
		return
	}

	pwd := &passwordResetModel.PasswordReset{Email: email}
	if err := pwd.Create(); err != nil {
		flash.NewDangerFlash(c, "重置密码链接生成失败: "+err.Error())
		controllers.RedirectRouter(c, "password.request")
		return
	}

	// 发送邮件
	if err := helpers.SendResetPasswordEmail(pwd); err != nil {
		flash.NewDangerFlash(c, "重置密码邮件发送失败: "+err.Error())
		passwordResetModel.DeleteByEmail(pwd.Email) // 删除 token
	} else {
		flash.NewSuccessFlash(c, "重置密码已发送到你的邮箱上，请注意查收。")
	}

	controllers.RedirectRouter(c, "password.request")
}

// ShowResetForm 重置密码页面
func ShowResetForm(c *gin.Context) {
	token := c.Param("token")
	p, err := passwordResetModel.GetByToken(token)
	if err != nil {
		controllers.Render404(c)
		return
	}

	controllers.Render(c, "auth/password/reset", gin.H{
		"token": token,
		"email": p.Email,
	})
}

// Reset 重置密码
func Reset(c *gin.Context) {
	token := c.PostForm("token")
	form := passowordRequest.PassWordResetForm{
		Token:                token,
		Password:             c.PostForm("password"),
		PasswordConfirmation: c.PostForm("password_confirmation"),
	}
	user, ok := form.ValidateAndUpdate(c)
	if !ok {
		controllers.RedirectRouter(c, "password.reset", "token", token)
		return
	}

	auth.Login(c, user)
	flash.NewSuccessFlash(c, "密码更新成功，您已成功登录！")
	controllers.RedirectRouter(c, "root")
}
