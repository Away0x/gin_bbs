package authorization

import (
	"gin_bbs/config"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils/socials/weixin"
	"gin_bbs/pkg/ginutils/utils"
	"gin_bbs/pkg/ginutils/validate"

	"github.com/lexkong/log"
)

var (
	types = []string{"weixin"}
)

// Social -
// 要么传 OpenID 和 AccessToken，要么传 Code
/*
{ "access_token": "xxx", "openid": "xxx" } 或者 { "code": "xxx" }
*/
type Social struct {
	validate.Validate
	SocialType  string `json:"-"`
	Code        string `json:"code"`
	AccessToken string `json:"access_token"`
	OpenID      string `json:"openid"`
}

func (s *Social) socialTypeValidator() validate.ValidatorFunc {
	return func() string {
		if !utils.InStringSlice(types, s.SocialType) {
			return "social_type 错误"
		}

		return ""
	}
}

// RegisterValidators 注册验证器
func (s *Social) RegisterValidators() validate.ValidatorMap {
	return validate.ValidatorMap{
		"social_type": {
			s.socialTypeValidator(),
		},
		"code": {
			func() string {
				// s.AccessToken 和 s.Code 是互斥关系
				if s.AccessToken != "" && s.Code != "" {
					return "code 传参错误"
				}

				return ""
			},
		},
		"access_token": {
			func() string {
				// s.AccessToken 和 s.Code 是互斥关系
				if s.AccessToken != "" && s.Code != "" {
					return "access_token 传参错误"
				}

				return ""
			},
		},
		"openid": {
			func() string {
				if s.SocialType == "weixin" && s.Code == "" {
					if s.OpenID == "" {
						return "openid 传参错误"
					}
				}

				return ""
			},
		},
	}
}

// Run -
func (s *Social) Run() (*weixin.WeixinUserInfo, *errno.Errno) {
	if config.AppConfig.WeixinAppID == "" || config.AppConfig.WeixinAppSecret == "" {
		log.Warn("weixin config error: 未配置 WEIXIN CONFIG，请检查 config.yaml 配置")
		return nil, errno.New(errno.InternalServerError, "weixin config error: 未配置 WEIXIN CONFIG，请检查 config.yaml 配置")
	}
	var (
		err         error
		accessToken string
		openid      string
		userInfo    *weixin.WeixinUserInfo
	)

	if ok, _, errMap := validate.Run(s); !ok {
		return nil, errno.New(errno.ParamsError, errMap)
	}

	if s.Code != "" {
		// 获取 accessToken
		accessToken, openid, err = weixin.GetAccessToken(config.AppConfig.WeixinAppID, config.AppConfig.WeixinAppSecret, s.Code)
		if err != nil {
			return nil, errno.New(errno.SocialAuthorizationError, err)
		}
	} else {
		accessToken = s.AccessToken
		openid = s.OpenID
	}

	userInfo, err = weixin.UserFromToken(accessToken, openid)
	if err != nil {
		return nil, errno.New(errno.SocialAuthorizationError, err)
	}

	return userInfo, nil
}
