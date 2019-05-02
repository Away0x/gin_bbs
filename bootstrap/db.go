package bootstrap

import (
	"errors"
	"fmt"
	"gin_bbs/app/models/user"
	"gin_bbs/config"
	"gin_bbs/database"

	"github.com/jinzhu/gorm"
)

// SetupDB db setup
func SetupDB(needMock bool) (*gorm.DB, error) {
	db := database.InitDB()

	// db migrate
	db.AutoMigrate(
		&user.User{},
	)

	// mock data
	if do := factoryMake(needMock); do {
		return db, errors.New("mock data")
	}

	return db, nil
}

// 数据 mock
func factoryMake(needMock bool) (do bool) {
	// 只有非 release 时才可用该函数
	if config.AppConfig.RunMode == config.RunmodeRelease || !needMock {
		return false
	}

	fmt.Print("\n\n-------------------------------------------------- MOCK --------------------------------------------------\n\n")
	// factory function
	return true
}
