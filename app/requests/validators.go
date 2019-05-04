package requests

import (
	userModel "gin_bbs/app/models/user"
	"gin_bbs/pkg/ginutils/captcha"
	"gin_bbs/pkg/ginutils/validate"
)

// EmailUniqueValidator 邮箱唯一
func EmailUniqueValidator(email string) validate.ValidatorFunc {
	return func() (msg string) {
		if _, err := userModel.GetByEmail(email); err != nil {
			return ""
		}
		return "邮箱已经被注册过了"
	}
}

// CaptchaValidator 验证码验证
func CaptchaValidator(captchaID, captchaVal string) validate.ValidatorFunc {
	return func() (msg string) {
		if ok := captcha.Verify(captchaID, captchaVal); ok {
			return ""
		}
		return "验证码错误"
	}
}
