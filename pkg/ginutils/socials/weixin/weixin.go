package weixin

import (
	"errors"
	"fmt"
	"gin_bbs/pkg/ginutils/utils"

	"encoding/json"

	"github.com/buger/jsonparser"
)

// func getOAuthURL(appid string, redirectURL string, responseType string, scope string, state string) string {
//   return fmt.Sprintf(
//     "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=%s&scope=%s&state=%s#wechat_redirect",
//     appid,
//     redirectURL,
//     responseType,
//     scope,
//     state,
//   )
// }

// WeixinUserInfo 微信用户信息
type WeixinUserInfo struct {
	OpenID     string      `json:"openid"`
	NickName   string      `json:"nickname"`
	Sex        int         `json:"sex"`
	Language   string      `json:"language"`
	City       string      `json:"city"`
	Province   string      `json:"province"`
	Country    string      `json:"country"`
	HeadImgURL string      `json:"headimgurl"`
	Privilege  interface{} `json:"privilege"`
	Unionid    string      `json:"unionid"`
}

// GetAccessToken 获取微信 access_token
func GetAccessToken(appid, secret, code string) (accessToken, openid string, err error) {
	api := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		appid,
		secret,
		code,
	)

	body, err := utils.RequestGet(api)
	if err != nil {
		return "", "", err
	}

	accessToken, err = jsonparser.GetString(body, "access_token")
	openid, err = jsonparser.GetString(body, "openid")
	if err != nil {
		return "", "", errors.New(string(body))
	}

	return accessToken, openid, nil
}

// UserFromToken 获取微信用户信息
func UserFromToken(accessToken string, openID string) (*WeixinUserInfo, error) {
	api := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN",
		accessToken,
		openID,
	)

	body, err := utils.RequestGet(api)
	if err != nil {
		return nil, err
	}

	var info *WeixinUserInfo
	if err = json.Unmarshal(body, &info); err != nil {
		return nil, err
	}

	if info.OpenID == "" {
		return nil, errors.New(string(body))
	}

	return info, nil
}
