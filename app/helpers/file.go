package helpers

import (
	"errors"
	"gin_bbs/config"
	"gin_bbs/pkg/file"
	"gin_bbs/pkg/ginutils/utils"
	"mime/multipart"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	UploadFolder = "/uploads/"
	// 图片存储路径
	ImagesUploadFolder = UploadFolder + "images/"
)

// 获取静态文件访问路径
func GetPublicPath(path string) string {
	return config.AppConfig.PublicPath + path
}

// 生成以时间(年月/日)作为文件夹名
func CreateBaseTimeFolderName() string {
	now := time.Now()
	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	if len(month) == 1 {
		month = "0" + month
	}
	day := strconv.Itoa(now.Day())

	return "/" + year + month + "/" + day
}

// 生成随机文件名
// file_prefix 文件名前缀
// defaultExt 默认 ext (有些文件上传时没带文件后缀)
func CreateRandomFileName(f *multipart.FileHeader, file_prefix, defaultExt string) (string, string) {
	fileName := f.Filename
	// 获取文件的后缀名，因图片从剪贴板里黏贴时后缀名为空，所以此处确保后缀一直存在
	ext := path.Ext(fileName)
	if ext == "" && defaultExt != "" {
		ext = defaultExt // 默认的后缀
	}

	// 拼接文件名，加前缀是为了增加辨析度，前缀可以是相关数据模型的 ID
	// 值如：1_1493521050_7BVc9v9ujP.png
	fileName = file_prefix + "_" + time.Now().String() + "_" + string(utils.RandomCreateBytes(10)) + ext

	return fileName, ext
}

// 保存图片
func SaveImage(f *multipart.FileHeader, folderName, file_prefix string) (string, error) {
	fileName, ext := CreateRandomFileName(f, file_prefix, ".png")
	fullPath := GetPublicPath(ImagesUploadFolder + folderName + CreateBaseTimeFolderName())

	ext = strings.ToLower(ext)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" && ext != ".bmp" && ext != ".gif" {
		return "", errors.New("文件格式错误，不能上传 " + ext + "格式的文件")
	}

	if err := file.SaveFile(f, fullPath, fileName); err != nil {
		return "", err
	}

	return config.AppConfig.URL + "/" + fullPath + "/" + fileName, nil
}
