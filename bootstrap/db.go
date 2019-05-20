package bootstrap

import (
	"errors"
	"fmt"
	"gin_bbs/app/models/category"
	"gin_bbs/app/models/image"
	"gin_bbs/app/models/notification"
	passwordreset "gin_bbs/app/models/password_reset"
	"gin_bbs/app/models/permission"
	"gin_bbs/app/models/reply"
	"gin_bbs/app/models/topic"
	"gin_bbs/app/models/user"
	"gin_bbs/config"
	"gin_bbs/database"
	"gin_bbs/database/factory"

	"github.com/jinzhu/gorm"
)

// SetupDB db setup
func SetupDB(needMock bool) (*gorm.DB, error) {
	db := database.InitDB()

	// db migrate
	db.AutoMigrate(
		// permission
		&permission.Permission{},
		&permission.Role{},
		&permission.ModelHasPermission{},
		&permission.ModelHasRole{},
		&permission.RoleHasPermission{},

		&user.User{},
		&passwordreset.PasswordReset{},
		&category.Category{},
		&topic.Topic{},
		&reply.Reply{},
		&notification.Notification{},
		&image.Image{},
	)

	// 只有非 release 时才可 mock
	if needMock {
		if config.AppConfig.RunMode == config.RunmodeRelease {
			panic("[mock] 请在非生产环境中使用")
		}

		fmt.Print("\n\n-------------------------------------------------- MOCK --------------------------------------------------\n\n")
		factory.Mock()
		return db, errors.New("mock data")
	}

	return db, nil
}
