package helpers

import (
	"gin_bbs/config"
	"gin_bbs/pkg/mail"
	"gin_bbs/pkg/utils"
)

// SendMail 发送邮件
func SendMail(mailTo []string, subject string, templatePath string, tplData map[string]interface{}) error {
	filePath := config.AppConfig.ViewsPath + "/" + templatePath
	body, err := utils.ReadTemplateToString(templatePath, filePath, tplData)
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
