package user

import (
	userModel "gin_bbs/app/models/user"
	"gin_bbs/pkg/ginutils/validate"
)

// 以后可以改为 tag 来调用验证器函数
type UserCreateForm struct {
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

// Validate : 验证函数
func (u *UserCreateForm) Validate() (errors []string) {
	errors = validate.RunValidators(
		validate.ValidatorMap{
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
		},
		validate.ValidatorMsgArr{
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
		},
	)

	return errors
}

// ValidateAndSave 验证参数并且创建用户
func (u *UserCreateForm) ValidateAndSave() (user *userModel.User, errors []string) {
	errors = u.Validate()

	if len(errors) != 0 {
		return nil, errors
	}

	// 创建用户
	user = &userModel.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}

	if err := user.Create(); err != nil {
		errors = append(errors, "用户创建失败: "+err.Error())
		return nil, errors
	}

	return user, []string{}
}
