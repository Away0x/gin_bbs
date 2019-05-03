package user

import (
	userModel "gin_bbs/app/models/user"

	"gin_bbs/pkg/ginutils/validate"

	"github.com/gin-gonic/gin"
)

type UserLoginForm struct {
	validate.Validate
	Email    string
	Password string
}

func (u *UserLoginForm) RegisterValidators() validate.ValidatorMap {
	return validate.ValidatorMap{
		"email": {
			validate.RequiredValidator(u.Email),
			validate.MaxLengthValidator(u.Email, 255),
			validate.EmailValidator(u.Email),
		},
		"password": {
			validate.RequiredValidator(u.Password),
		},
	}
}

func (*UserLoginForm) RegisterMessages() validate.MessagesMap {
	return validate.MessagesMap{
		"email": {
			"邮箱不能为空",
			"邮箱长度不能大于 255 个字符",
			"邮箱格式错误",
		},
		"password": {
			"密码不能为空",
		},
	}
}

// ValidateAndLogin 验证参数并且获取用户
func (u *UserLoginForm) ValidateAndGetUser(c *gin.Context) (*userModel.User, []string) {
	ok, errArr, _ := validate.Run(u)

	if !ok {
		return nil, errArr
	}

	// 通过邮箱获取用户，并且判断密码是否正确
	user, err := userModel.GetByEmail(u.Email)
	if err != nil {
		errArr = append(errArr, "该邮箱没有注册过用户: "+err.Error())
		return nil, errArr
	}

	if err := user.Compare(u.Password); err != nil {
		errArr = append(errArr, "很抱歉，您的邮箱和密码不匹配: "+err.Error())
		return nil, errArr
	}

	return user, errArr
}
