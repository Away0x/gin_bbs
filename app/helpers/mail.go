package helpers

import (
	"gin_bbs/config"
	"gin_bbs/pkg/ginutils/mail"
	"gin_bbs/pkg/ginutils/router"
	"path"

	passwordResetModel "gin_bbs/app/models/password_reset"
	userModel "gin_bbs/app/models/user"

	"github.com/flosch/pongo2"
)

// SendMail 发送邮件
func SendMail(mailTo []string, subject string, templateName string, tplData map[string]interface{}) error {
	filename := path.Join(config.AppConfig.ViewsPath, templateName)
	template := pongo2.Must(pongo2.FromCache(filename))

	body, err := template.Execute(tplData)
	if err != nil {
		return err
	}

	mail := &mail.Mail{
		Driver:   config.MailConfig.Driver,
		Host:     config.MailConfig.Host,
		Port:     config.MailConfig.Port,
		User:     config.MailConfig.User,
		Password: config.MailConfig.Password,
		FromName: config.MailConfig.FromName,
		MailTo:   mailTo,
		Subject:  subject,
		Body:     body,
	}

	return mail.Send()
}

// SendVerifyEmail 发送激活用户的邮件
func SendVerifyEmail(u *userModel.User) error {
	subject := "感谢注册 Weibo 应用！请确认你的邮箱。"
	tpl := "mail/verify.html"
	verifyURL := router.G("verification.verify", "token", u.ActivationToken)

	return SendMail([]string{u.Email}, subject, tpl, map[string]interface{}{"URL": verifyURL})
}

// SendResetPasswordEmail 发送重置密码邮件
func SendResetPasswordEmail(pwd *passwordResetModel.PasswordReset) error {
	subject := "重置密码！请确认你的邮箱。"
	tpl := "mail/reset_password.html"
	resetPasswordURL := router.G("password.reset", "token", pwd.Token)

	return SendMail([]string{pwd.Email}, subject, tpl, map[string]interface{}{"URL": resetPasswordURL})
}
