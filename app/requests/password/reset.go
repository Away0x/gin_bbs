package passsword

import (
	passwordResetModel "gin_bbs/app/models/password_reset"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/pkg/ginutils/flash"
	"gin_bbs/pkg/ginutils/validate"

	"github.com/gin-gonic/gin"
)

type PassWordResetForm struct {
	validate.Validate
	Email                string
	Token                string
	Password             string
	PasswordConfirmation string
}

func (p *PassWordResetForm) tokenExistValidator() validate.ValidatorFunc {
	return func() (msg string) {
		if m, err := passwordResetModel.GetByToken(p.Token); err == nil {
			p.Email = m.Email
			return ""
		}
		return "该 token 不存在"
	}
}

// RegisterValidators 注册验证器
func (p *PassWordResetForm) RegisterValidators() validate.ValidatorMap {
	return validate.ValidatorMap{
		"password": {
			validate.RequiredValidator(p.Password),
			validate.MixLengthValidator(p.Password, 6),
			validate.EqualValidator(p.Password, p.PasswordConfirmation),
		},
		"token": {
			validate.RequiredValidator(p.Token),
			p.tokenExistValidator(),
		},
	}
}

// RegisterMessages 注册错误信息
func (*PassWordResetForm) RegisterMessages() validate.MessagesMap {
	return validate.MessagesMap{
		"password": {
			"密码不能为空",
			"密码长度不能小于 6 个字符",
			"两次输入的密码不一致",
		},
		"token": {
			"token 不能为空",
			"该 token 不存在",
		},
	}
}

// ValidateAndUpdate 验证参数并且创建验证 pwd 的 token
func (p *PassWordResetForm) ValidateAndUpdate(c *gin.Context) (*userModel.User, bool) {
	ok, errArr, errMap := validate.Run(p)
	if !ok {
		validate.SaveValidateMessage(c, errArr, errMap)
		return nil, false
	}

	// 验证成功，删除 token
	if err := passwordResetModel.DeleteByToken(p.Token); err != nil {
		flash.NewDangerFlash(c, "重置密码失败: "+err.Error())
		return nil, false
	}

	// 更新用户密码
	user, err := userModel.GetByEmail(p.Email)
	if err != nil {
		flash.NewDangerFlash(c, "重置密码失败: "+err.Error())
		return nil, false
	}
	user.Password = p.Password
	if err = user.Update(); err != nil {
		flash.NewDangerFlash(c, "重置密码失败: "+err.Error())
		return nil, false
	}

	return user, true
}
