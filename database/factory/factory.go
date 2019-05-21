package factory

import (
	"gin_bbs/database"
)

// dropAndCreateTable - 清空表
func dropAndCreateTable(table interface{}) {
	database.DB.DropTable(table)
	database.DB.CreateTable(table)
}

// Mock -
func Mock() {
	PermissionTableSeeder(true)
	UsersTableSeeder(true)
	CategoryTableSeeder(true)
	TopicTableSeeder(true)
	ReplyTableSeeder(true)
	LinksTableSeeder(true)
}
