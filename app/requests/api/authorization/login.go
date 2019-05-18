package authorization

import (
	userModel "gin_bbs/app/models/user"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils/validate"
)

// Login -
type Login struct {
	validate.Validate
	Username string `json:"username"` // 用户可以使用邮箱或者手机号登录
	Password string `json:"password"`

	userFindKey string // 用于查找用户的 key 值 (phone or email)
}

func (l *Login) phoneOrEmailValidator() validate.ValidatorFunc {
	return func() string {
		msg := validate.EmailValidator(l.Username)()
		if msg != "" {
			msg := validate.PhoneValidator(l.Username)()
			if msg != "" {
				return "用户名请输入正确的手机号或者邮箱地址"
			}
			l.userFindKey = "phone"
			return ""
		}

		l.userFindKey = "email"
		return ""
	}
}

// RegisterValidators 注册验证器
func (l *Login) RegisterValidators() validate.ValidatorMap {
	return validate.ValidatorMap{
		"username": {
			validate.RequiredValidator(l.Username),
			l.phoneOrEmailValidator(),
		},
		"password": {
			validate.RequiredValidator(l.Password),
			validate.MinLengthValidator(l.Password, 6),
		},
	}
}

// RegisterMessages 注册错误信息
func (*Login) RegisterMessages() validate.MessagesMap {
	return validate.MessagesMap{
		"username": {
			"用户名不能为空",
		},
		"password": {
			"密码不能为空",
			"密码长度不能小于 6 个字符",
		},
	}
}

// Run -
func (l *Login) Run() (*userModel.User, *errno.Errno) {
	ok, _, errMap := validate.Run(l)
	if !ok {
		return nil, errno.New(errno.ParamsError, errMap)
	}

	// 用户可以使用邮箱或者手机号登录
	// fmt.Println(l.userFindKey, l.Username)
	user, err := userModel.First(l.userFindKey+"= ?", l.Username)
	if err != nil {
		return nil, errno.New(errno.DatabaseError, "没有找到该用户")
	}

	// 密码验证
	if err = user.Compare(l.Password); err != nil {
		return nil, errno.New(errno.LoginError, err)
	}

	return user, nil
}
