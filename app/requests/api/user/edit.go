package user

import (
	imageModel "gin_bbs/app/models/image"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/app/requests"
	"gin_bbs/pkg/constants"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils/validate"
)

// Edit -
type Edit struct {
	validate.Validate
	UserID        uint   `json:"-"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Introduction  string `json:"introduction"`
	AvatarImageID int    `json:"avatar_image_id"`
}

func (e *Edit) imageIDValidator() validate.ValidatorFunc {
	return func() string {
		if e.AvatarImageID == 0 {
			return ""
		}
		i, err := imageModel.Get(int(e.AvatarImageID))
		if err != nil {
			return ""
		}

		if i.Type != imageModel.TypeAvatar || i.UserID != e.UserID {
			return "图片不存在"
		}

		return ""
	}
}

// IsStrict 有错误即退出
func (*Edit) IsStrict() bool {
	return true
}

// RegisterValidators 注册验证器
func (e *Edit) RegisterValidators() validate.ValidatorMap {
	return validate.ValidatorMap{
		"name": {
			validate.RequiredValidator(e.Name),
			validate.BetweenValidator(e.Name, 3, 25),
			validate.RegexpValidator(e.Name, constants.UserNameRegex),
			requests.NameUniqueValidator(e.Name, int(e.UserID)),
		},
		"email": {
			validate.EmailValidator(e.Email),
		},
		"introduction": {
			validate.MaxLengthValidator(e.Email, 80),
		},
		"avatar_image_id": {
			e.imageIDValidator(),
		},
	}
}

// RegisterMessages 注册错误信息
func (*Edit) RegisterMessages() validate.MessagesMap {
	return validate.MessagesMap{
		"name": {
			"用户名不能为空。",
			"用户名必须介于 3 - 25 个字符之间。",
			"用户名只支持英文、数字、横杆和下划线。",
			"用户名已被占用，请重新填写",
		},
	}
}

// Run -
func (e *Edit) Run(user *userModel.User) *errno.Errno {
	ok, _, errMap := validate.Run(e)
	if !ok {
		return errno.New(errno.ParamsError, errMap)
	}

	user.Name = e.Name
	user.Introduction = e.Introduction
	if e.Email != "" {
		user.Email = e.Email
	}
	if e.AvatarImageID != 0 {
		img, err := imageModel.Get(e.AvatarImageID)
		if err == nil {
			user.Avatar = img.Path
		}
	}

	if err := user.Update(); err != nil {
		return errno.New(errno.DatabaseError, err)
	}

	return nil
}
