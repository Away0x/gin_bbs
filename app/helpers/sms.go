package helpers

import (
	"fmt"
	"gin_bbs/config"

	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
)

// SendSms 发送短信
func SendSms(phone string, code string) *ypclnt.Result {
	apiKey := config.AppConfig.YunPianApiKey
	if apiKey == "" {
		// debug 模式并且没配置云片 apiKey 不发送短信
		if config.AppConfig.RunMode != config.RunmodeRelease {
			return &ypclnt.Result{
				Code:   200,
				Msg:    "云片 apikey 未配置，请检查 config.yaml",
				Detail: fmt.Sprintf("【debug】phone: %s, code: %s", phone, code),
			}
		}

		return &ypclnt.Result{
			Code: -1,
			Msg:  "发送失败",
		}
	}

	clnt := ypclnt.New(apiKey)
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = phone
	param[ypclnt.TEXT] = "【" + config.AppConfig.Name + "】您的验证码是" + code

	return clnt.Sms().SingleSend(param)
}
