package password

import (
	userModel "gin_bbs/app/models/user"
	"gin_bbs/pkg/ginutils/validate"
)

func emailExistValidator(email string) validate.ValidatorFunc {
	return func() string {
		if _, err := userModel.GetByEmail(email); err == nil {
			return ""
		}
		return "该邮箱不存在"
	}
}
