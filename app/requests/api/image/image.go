package image

import (
	"gin_bbs/app/helpers"
	imageModel "gin_bbs/app/models/image"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/pkg/constants"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils/validate"
	"mime/multipart"
)

// Upload -
type Upload struct {
	validate.Validate
	Type  string
	Image *multipart.FileHeader
}

// RegisterValidators 注册验证器
func (u *Upload) RegisterValidators() validate.ValidatorMap {
	vs := validate.ValidatorMap{
		"type": {
			validate.RequiredValidator(u.Type),
			validate.StringRangeValidator(u.Type, imageModel.Types),
		},
		"image": {
			func() string {
				if u.Image == nil {
					return "image 必须存在"
				}
				return ""
			},
			validate.MimetypeValidator(u.Image, constants.UploadImageMimetypes),
		},
	}

	if u.Type == imageModel.TypeAvatar {
		vs["image"] = append(vs["image"], validate.ImageDimensionsValidator(u.Image,
			validate.DimensionsOptions{MinWidth: 208, MinHeight: 208}))
	}

	return vs
}

// RegisterMessages 注册错误信息
func (*Upload) RegisterMessages() validate.MessagesMap {
	return validate.MessagesMap{
		"image": {
			"",
			"",
			"图片的清晰度不够，宽和高需要 200px 以上",
		},
	}
}

// Run -
func (u *Upload) Run(user *userModel.User) (*imageModel.Image, *errno.Errno) {
	ok, _, errMap := validate.Run(u)
	if !ok {
		return nil, errno.New(errno.ParamsError, errMap)
	}

	maxWidth := 362
	if u.Type != imageModel.TypeAvatar {
		maxWidth = 1024
	}

	path, err := helpers.SaveImage(u.Image, u.Type+"s", user.GetIDstring(), maxWidth)
	if err != nil {
		return nil, errno.New(errno.UploadError, err)
	}

	// 保存
	i := &imageModel.Image{
		Type:   u.Type,
		UserID: user.ID,
		Path:   path,
	}
	if err = i.Create(); err != nil {
		return nil, errno.New(errno.DatabaseError, err)
	}

	return i, nil
}
