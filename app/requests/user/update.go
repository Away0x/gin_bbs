package user

import (
	"gin_bbs/app/helpers"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/app/requests"
	"gin_bbs/pkg/constants"
	"gin_bbs/pkg/ginutils/flash"
	"gin_bbs/pkg/ginutils/validate"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

// UserUpdateForm -
type UserUpdateForm struct {
	validate.Validate
	ID           int
	Name         string
	Email        string
	Introduction string
	Avatar       *multipart.FileHeader
}

// IsStrict 有错误即退出
func (*UserUpdateForm) IsStrict() bool {
	return true
}

// RegisterValidators 注册验证器
func (u *UserUpdateForm) RegisterValidators() validate.ValidatorMap {
	return validate.ValidatorMap{
		"name": {
			validate.RequiredValidator(u.Name),
			validate.BetweenValidator(u.Name, 3, 25),
			validate.RegexpValidator(u.Name, constants.UserNameRegex),
			requests.NameUniqueValidator(u.Name, u.ID),
		},
		"email": {
			validate.RequiredValidator(u.Email),
			validate.MaxLengthValidator(u.Email, 255),
			validate.EmailValidator(u.Email),
		},
		"introduction": {
			validate.MaxLengthValidator(u.Introduction, 80),
		},
		"avatar": {
			validate.MimetypeValidator(u.Avatar, constants.UploadImageMimetypes),
			validate.ImageDimensionsValidator(u.Avatar, validate.DimensionsOptions{MinWidth: 208, MinHeight: 208}),
		},
	}
}

// RegisterMessages 错误信息
func (*UserUpdateForm) RegisterMessages() validate.MessagesMap {
	return validate.MessagesMap{
		"name": {
			"用户名不能为空。",
			"用户名必须介于 3 - 25 个字符之间",
			"用户名只支持英文、数字、横杠和下划线。",
			"用户名已被占用，请重新填写",
		},
		"introduction": {
			"用户介绍不得大于 80 个字",
		},
		"avatar": {
			"头像必须是 jpeg, bmp, png, gif 格式的图片",
			"图片的清晰度不够，宽和高需要 208px 以上",
		},
	}
}

// ValidateAndUpdate 验证参数并且更新用户
func (u *UserUpdateForm) ValidateAndUpdate(c *gin.Context, user *userModel.User) bool {
	ok, errArr, errMap := validate.Run(u)
	if !ok {
		validate.SaveValidateMessage(c, errArr, errMap)
		return false
	}

	// 更新用户
	user.Name = u.Name
	user.Email = u.Email
	user.Introduction = u.Introduction
	// 如果有上传用户头像
	if u.Avatar != nil {
		avatarPath, err := helpers.SaveImage(u.Avatar, "avatars", user.GetIDstring(), 416)
		if err != nil {
			validate.AddMessageAndSaveToFlash(c, "avatar", "头像上传失败: "+err.Error(), errArr, errMap)
			return false
		}
		user.Avatar = avatarPath
	}

	if err := user.Update(); err != nil {
		flash.NewDangerFlash(c, "用户更新失败: "+err.Error())
		return false
	}

	return true
}
