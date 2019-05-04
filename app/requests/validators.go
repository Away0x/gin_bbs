package requests

import (
	userModel "gin_bbs/app/models/user"
	"gin_bbs/pkg/ginutils/captcha"
	"gin_bbs/pkg/ginutils/validate"
)

// NameUniqueValidator name 唯一
func NameUniqueValidator(name string, id int) validate.ValidatorFunc {
	return func() (msg string) {
		u, err := userModel.GetByName(name)
		if err != nil || u.ID == uint(id) {
			return ""
		}

		return "用户名已经被注册过了"
	}
}

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
