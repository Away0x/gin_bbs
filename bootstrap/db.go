package bootstrap

import (
	"errors"
	"fmt"
	passwordreset "gin_bbs/app/models/password_reset"
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
		&user.User{},
		&passwordreset.PasswordReset{},
	)

	// 只有非 release 时才可 mock
	if config.AppConfig.RunMode != config.RunmodeRelease && needMock {
		fmt.Print("\n\n-------------------------------------------------- MOCK --------------------------------------------------\n\n")
		factory.Mock()
		return db, errors.New("mock data")
	}

	return db, nil
}
