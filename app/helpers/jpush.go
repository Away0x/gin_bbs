package helpers

import (
	"gin_bbs/config"

	jpushclient "github.com/ylywyn/jpush-api-go-client"
)

var (
	// JpushClient -
	JpushClient *JPush
)

// JPush 极光推送
type JPush struct {
	Key    string
	Secret string
}

// NewJPush -
func NewJPush() *JPush {
	return &JPush{
		Key:    config.AppConfig.JPushKey,
		Secret: config.AppConfig.JPushSecret,
	}
}

// Send -
func (j *JPush) Send(content string, pushids []string) {
	var pf jpushclient.Platform
	pf.All()

	var ad jpushclient.Audience
	ad.SetID(pushids)

	var msg jpushclient.Message
	msg.Content = content

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetMessage(&msg)

	bytes, _ := payload.ToBytes()
	c := jpushclient.NewPushClient(j.Secret, j.Key)
	c.Send(bytes)
}
