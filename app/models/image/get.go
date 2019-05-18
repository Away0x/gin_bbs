package image

import (
	"gin_bbs/app/models/user"
	"gin_bbs/database"
)

// Get -
func Get(id int) (*Image, error) {
	i := &Image{}
	if err := database.DB.First(&i, id).Error; err != nil {
		return nil, err
	}

	return i, nil
}

// User 获取 user
func (i *Image) User() (u *user.User, err error) {
	u, err = user.Get(int(i.UserID))
	if err != nil {
		return nil, err
	}

	return u, err
}
