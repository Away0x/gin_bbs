package user

import (
	userModel "gin_bbs/app/models/user"
	"gin_bbs/pkg/ginutils/validate"

	"github.com/gin-gonic/gin"
)

// 以后可以改为 tag 来调用验证器函数
type UserCreateForm struct {
	validate.Validate
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string
}

func (u *UserCreateForm) emailUniqueValidator() validate.ValidatorFunc {
	return func() (msg string) {
		if _, err := userModel.GetByEmail(u.Email); err != nil {
			return ""
		}
		return "邮箱已经被注册过了"
	}
}

// RegisterValidators 注册验证器
func (u *UserCreateForm) RegisterValidators() validate.ValidatorMap {
	return validate.ValidatorMap{
		"name": {
			validate.RequiredValidator(u.Name),
			validate.MaxLengthValidator(u.Name, 50),
		},
		"email": {
			validate.RequiredValidator(u.Email),
			validate.MaxLengthValidator(u.Email, 255),
			validate.EmailValidator(u.Email),
			u.emailUniqueValidator(),
		},
		"password": {
			validate.RequiredValidator(u.Password),
			validate.MixLengthValidator(u.Password, 6),
			validate.EqualValidator(u.Password, u.PasswordConfirmation),
		},
	}
}

// RegisterMessages 注册错误信息
func (*UserCreateForm) RegisterMessages() validate.MessagesMap {
	return validate.MessagesMap{
		"name": {
			"名称不能为空",
			"名称长度不能大于 50 个字符",
		},
		"email": {
			"邮箱不能为空",
			"邮箱长度不能大于 255 个字符",
			"邮箱格式错误",
			"邮箱已经被注册过了",
		},
		"password": {
			"密码不能为空",
			"密码长度不能小于 6 个字符",
			"两次输入的密码不一致",
		},
	}
}

// ValidateAndSave 验证参数并且创建用户
func (u *UserCreateForm) ValidateAndSave(c *gin.Context) (bool, *userModel.User) {
	ok, errArr, errMap := validate.Run(u)

	if !ok {
		validate.SaveValidateMessage(c, errArr, errMap)
		return false, nil
	}

	// 创建用户
	user := &userModel.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}

	if err := user.Create(); err != nil {
		errMap["other"] = make([]string, 0)
		errMap["other"] = append(errMap["other"], "用户创建失败: "+err.Error())
		validate.SaveValidateMessage(c, errArr, errMap)
		return false, nil
	}

	return true, user
}
