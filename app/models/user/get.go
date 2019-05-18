package user

import (
	"gin_bbs/app/models"
	"gin_bbs/database"
)

// Get -
func Get(id int) (*User, error) {
	if cachedUser, ok := getFromCache(id); ok {
		return cachedUser, nil
	}

	user := &User{}
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	setToCache(user)

	return user, nil
}

// First -
func First(query interface{}, args ...interface{}) (*User, error) {
	user := &User{}
	d := database.DB.Where(query, args...).First(&user)
	return user, d.Error
}

// GetByName -
func GetByName(name string) (*User, error) {
	user := &User{}
	d := database.DB.Where("name = ?", name).First(&user)
	return user, d.Error
}

// GetByEmail -
func GetByEmail(email string) (*User, error) {
	user := &User{}
	d := database.DB.Where("email = ?", email).First(&user)
	return user, d.Error
}

// GetByPhone -
func GetByPhone(phone string) (*User, error) {
	user := &User{}
	d := database.DB.Where("phone = ?", phone).First(&user)
	return user, d.Error
}

// GetByActivationToken -
func GetByActivationToken(token string) (*User, error) {
	user := &User{}
	d := database.DB.Where("activation_token = ?", token).First(&user)
	return user, d.Error
}

// GetByRememberToken -
func GetByRememberToken(token string) (*User, error) {
	user := &User{}
	d := database.DB.Where("remember_token = ?", token).First(&user)
	return user, d.Error
}

// GetByWeixinUnionID -
func GetByWeixinUnionID(unionid string) (*User, error) {
	user := &User{}
	d := database.DB.Where("weixin_unionid = ?", unionid).First(&user)
	return user, d.Error
}

// GetByWeixinOpenID -
func GetByWeixinOpenID(openid string) (*User, error) {
	user := &User{}
	d := database.DB.Where("weixin_openid = ?", openid).First(&user)
	return user, d.Error
}

// List 获取用户列表
func List(offset, limit int) (users []*User, err error) {
	users = make([]*User, 0)

	if err := database.DB.Offset(offset).Limit(limit).Order("id").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

// All -
func All() (users []*User, err error) {
	users = make([]*User, 0)

	if err := database.DB.Order("id").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

// AllID -
func AllID() (ids []uint, err error) {
	ids = make([]uint, 0)
	users, err := All()
	if err != nil {
		return ids, err
	}

	for _, u := range users {
		ids = append(ids, u.ID)
	}

	return ids, nil
}

// AllCount 总用户数
func AllCount() (count int, err error) {
	err = database.DB.Model(&User{}).Count(&count).Error
	return
}

// IsActivated 是否已激活
func (u *User) IsActivated() bool {
	return u.Activated == models.TrueTinyint
}
