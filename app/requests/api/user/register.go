package user

import (
	"gin_bbs/app/cache"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/app/requests"
	"gin_bbs/pkg/constants"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils/validate"
)

// Register -
type Register struct {
	validate.Validate
	Name             string `json:"name"`
	Password         string `json:"password"`
	VerificationKey  string `json:"verification_key"`
	VerificationCode string `json:"verification_code"`
}

// RegisterValidators 注册验证器
func (u *Register) RegisterValidators() validate.ValidatorMap {
	return validate.ValidatorMap{
		"name": {
			validate.RequiredValidator(u.Name),
			validate.BetweenValidator(u.Name, 3, 25),
			validate.RegexpValidator(u.Name, constants.UserNameRegex),
			requests.NameUniqueValidator(u.Name, 0),
		},
		"password": {
			validate.RequiredValidator(u.Password),
			validate.MinLengthValidator(u.Password, 6),
		},
		"verification_key": {
			validate.RequiredValidator(u.VerificationKey),
		},
		"verification_code": {
			validate.RequiredValidator(u.VerificationCode),
		},
	}
}

// RegisterMessages 注册错误信息
func (*Register) RegisterMessages() validate.MessagesMap {
	return validate.MessagesMap{
		"name": {
			"用户名不能为空",
			"用户名长度应在 3 ~ 25 个字符之间",
			"用户名不合法",
			"该用户名已被注册",
		},
		"password": {
			"密码不能为空",
			"密码长度不能小于 6 个字符",
		},
	}
}

// ValidateAndCreateUser -
func (u *Register) ValidateAndCreateUser() *errno.Errno {
	ok, _, errMap := validate.Run(u)
	if !ok {
		return errno.New(errno.ParamsError, errMap)
	}

	cachedData, ok := cache.GetStringMap(u.VerificationKey)
	if !ok {
		return errno.New(errno.ParamsError, map[string]string{"verification_code": "验证码已失效"})
	}

	if cachedData["code"] != u.VerificationCode {
		return errno.New(errno.ParamsError, map[string]string{"verification_code": "验证码错误"})
	}

	// 创建用户
	user := &userModel.User{
		Name:     u.Name,
		Phone:    cachedData["phone"],
		Password: u.Password,
	}
	if err := user.Create(); err != nil {
		return errno.New(errno.DatabaseError, err)
	}

	// 清缓存
	cache.Del(u.VerificationKey)

	return nil
}
