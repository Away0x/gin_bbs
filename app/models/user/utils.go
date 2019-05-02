package user

import (
	"gin_bbs/pkg/ginutils/utils"
)

// Encrypt 对密码进行加密
func (u *User) Encrypt() (err error) {
	u.Password, err = utils.Encrypt(u.Password)
	return
}

// Compare 验证用户密码
func (u *User) Compare(pwd string) (err error) {
	err = utils.Compare(u.Password, pwd)
	return
}
