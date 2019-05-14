package helpers

import (
	"gin_bbs/config"
	"strconv"

	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
)

// SendSms 发送短信
func SendSms(phone string, code int) *ypclnt.Result {
	apiKey := config.AppConfig.YunPianApiKey
	if apiKey == "" {
		return nil
	}

	clnt := ypclnt.New(apiKey)
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = phone
	param[ypclnt.TEXT] = "【" + config.AppConfig.Name + "】您的验证码是" + strconv.Itoa(code)

	return clnt.Sms().SingleSend(param)
}
