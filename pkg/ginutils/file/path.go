package file

import (
	"gin_bbs/pkg/ginutils"
)

// PublicPath 相对于项目 public 静态文件目录的地址
func PublicPath(staticFilePath string) string {
	if string(staticFilePath[0]) == "/" {
		return ginutils.GetGinUtilsConfig().PublicPath + staticFilePath
	}
	return ginutils.GetGinUtilsConfig().PublicPath + "/" + staticFilePath
}

// StaticPath 生成项目静态文件地址
func StaticPath(staticFilePath string) string {
	return "/" + PublicPath(staticFilePath)
}
