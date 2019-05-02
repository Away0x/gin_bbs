package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"

	"gin_bbs/config"
)

/*
ERROR: Incorrect datetime value: '0000-00-00 00:00:00' for column 'hr519799901' at row 1
解决: 把 NO_ZERO_IN_DATE,NO_ZERO_DATE 去掉，然后重新设置
*/

// DB gorm
var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	db, err := gorm.Open(config.DBConfig.Connection, config.DBConfig.URL)
	if err != nil {
		log.Fatal("Database connection failed. Database url: "+config.DBConfig.URL+" error: ", err)
	} else {
		fmt.Print("\n\n------------------------------------------ GORM OPEN SUCCESS! -----------------------------------------------\n\n")
	}

	db.LogMode(config.DBConfig.Debug)
	DB = db

	return db
}
