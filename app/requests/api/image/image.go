package image

import (
	"gin_bbs/pkg/ginutils/validate"
	"mime/multipart"
)

// Upload -
type Upload struct {
	validate.Validate
	Type  string                `json:"type"`
	Image *multipart.FileHeader `json:"image"`
}

// RegisterValidators 注册验证器
func (u *Upload) RegisterValidators() validate.ValidatorMap {
	return validate.ValidatorMap{}
}

// RegisterMessages 注册错误信息
func (*Upload) RegisterMessages() validate.MessagesMap {
	return validate.MessagesMap{}
}
