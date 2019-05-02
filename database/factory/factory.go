package factory

import (
	"gin_bbs/database"
)

// DropAndCreateTable - 清空表
func DropAndCreateTable(table interface{}) {
	database.DB.DropTable(table)
	database.DB.CreateTable(table)
}
