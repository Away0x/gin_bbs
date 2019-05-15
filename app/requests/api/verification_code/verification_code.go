package vericode

import (
	"gin_bbs/app/cache"
	"gin_bbs/app/requests"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils/validate"
)

// VerificationCode -
type VerificationCode struct {
	validate.Validate
	CaptchaKey  string `json:"captcha_key"`
	CaptchaCode string `json:"captcha_code"`
}

// RegisterValidators 注册验证器
func (v *VerificationCode) RegisterValidators() validate.ValidatorMap {
	return validate.ValidatorMap{
		"captcha_key": {
			validate.RequiredValidator(v.CaptchaKey),
		},
		"captcha_code": {
			validate.RequiredValidator(v.CaptchaCode),
		},
	}
}

// Run -
func (v *VerificationCode) Run() (string, *errno.Errno) {
	ok, _, errMap := validate.Run(v)
	if !ok {
		return "", errno.New(errno.ParamsError, errMap)
	}

	// 从缓存中获取验证码 id
	cachedData, ok := cache.GetStringMap(v.CaptchaKey)
	if !ok {
		return "", errno.New(errno.ParamsError, map[string]string{"captcha_code": "验证码已失效"})
	}

	// 验证码验证
	if msg := requests.CaptchaValidator(cachedData["captcha_id"], v.CaptchaCode)(); msg != "" {
		cache.Del(v.CaptchaKey)
		return "", errno.New(errno.ParamsError, map[string]string{"captcha_code": msg})
	}

	cache.Del(v.CaptchaKey) // 清除验证码缓存
	return cachedData["phone"], nil
}
