package user

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"gin_bbs/app/models"
	"gin_bbs/pkg/ginutils/utils"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
)

const (
	// TableName 表名
	TableName = "users"
)

var (
	userCache = cache.New(30*time.Minute, 1*time.Hour)
)

// User 用户模型
type User struct {
	models.BaseModel
	Name         string `gorm:"column:name;type:varchar(255);not null" sql:"index"`
	Phone        string `gorm:"column:phone;type:varchar(255);unique;default:NULL" sql:"index"`
	Email        string `gorm:"column:email;type:varchar(255);unique;default:NULL" sql:"index"`
	Avatar       string `gorm:"column:avatar;type:varchar(255);not null"`
	Introduction string `gorm:"column:introduction;type:varchar(255);not null"`
	Password     string `gorm:"column:password;type:varchar(255)"` // 可使用微信登录，所以可以为空
	// 微信
	WeixinOpenID  string `gorm:"column:weixin_openid;type:varchar(255);unique;default:NULL"`
	WeixinUnionID string `gorm:"column:weixin_unionid;type:varchar(255);unique;default:NULL"` // 在用户将公众号绑定到微信开放平台帐号后，才会出现 unionid 字段
	// 用户激活
	ActivationToken string     `gorm:"column:activation_token;type:varchar(255)"`
	Activated       uint       `gorm:"column:activated;type:tinyint(1);not null"`
	EmailVerifiedAt *time.Time `gorm:"column:email_verified_at"` // 激活时间
	// 用户最后登录时间
	LastActivedAt *time.Time `gorm:"column:last_actived_at"`

	RememberToken     string `gorm:"column:remember_token;type:varchar(100)"`      // 用于实现记住我功能，存入 cookie 中，下次带上时，即可直接登录
	NotificationCount int    `gorm:"column:notification_count;not null;default:0"` // 未读通知数

	RegistrationID uint `gorm:"column:registration_id;unique;default:NULL"` // Jpush 中的唯一标识
}

// TableName 表名
func (User) TableName() string {
	return TableName
}

// BeforeCreate - hook
func (u *User) BeforeCreate() (err error) {
	if u.Password != "" {
		if isEncrypted := passwordEncrypted(u.Password); !isEncrypted {
			if err = u.Encrypt(); err != nil {
				return errors.New("User Model 创建失败: passwordEncrypted")
			}
		}
	}

	// 生成用户 remember_token
	if u.RememberToken == "" {
		u.RememberToken = string(utils.RandomCreateBytes(10))
	}

	// 生成用户激活 token
	if u.ActivationToken == "" {
		u.ActivationToken = string(utils.RandomCreateBytes(30))
	}

	// 生成用户头像
	if u.Avatar == "" {
		hash := md5.Sum([]byte(u.Email))
		u.Avatar = "http://www.gravatar.com/avatar/" + hex.EncodeToString(hash[:])
	}

	return err
}

// BeforeUpdate - hook
func (u *User) BeforeUpdate() (err error) {
	if isEncrypted := passwordEncrypted(u.Password); !isEncrypted {
		if err = u.Encrypt(); err != nil {
			return errors.New("User Model 更新失败")
		}
	}

	return
}

// BeforeDelete - hook
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	// 当用户删除时，删除其发布的话题
	tx.Exec("delete from topics where user_id = ?", u.ID)
	// 当用户删除时，删除其发布的回复
	tx.Exec("delete from replies where user_id = ?", u.ID)

	return
}

// ------------ private
func passwordEncrypted(pwd string) (status bool) {
	return len(pwd) == 60 // 长度等于 60 说明加密过了
}

func setToCache(user *User) {
	key := strconv.Itoa(int(user.ID))
	userCache.Set(key, user, cache.DefaultExpiration)
}

func getFromCache(id int) (*User, bool) {
	cachedUser, ok := userCache.Get(strconv.Itoa(id))
	if !ok {
		return nil, false
	}

	u, ok := cachedUser.(*User)
	if !ok {
		return nil, false
	}

	return u, true
}

func delCache(id int) {
	userCache.Delete(strconv.Itoa(id))
}
