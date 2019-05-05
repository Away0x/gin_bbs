package helpers

import (
	"errors"
	"gin_bbs/config"
	ginfile "gin_bbs/pkg/ginutils/file"
	"mime/multipart"
	"strings"
)

const (
	// UploadFolder 上传文件的文件夹路径
	UploadFolder = "/uploads/"
	// ImagesUploadFolder 图片存储路径
	ImagesUploadFolder = UploadFolder + "images/"
)

// SaveImage 保存图片
func SaveImage(f *multipart.FileHeader, folderName, filePrefix string) (string, error) {
	fileName, ext := ginfile.CreateRandomFileName(f, filePrefix, ".png")
	fullPath := ginfile.PublicPath(ImagesUploadFolder + folderName + ginfile.CreateBaseTimeFolderName())

	ext = strings.ToLower(ext)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" && ext != ".bmp" && ext != ".gif" {
		return "", errors.New("文件格式错误，不能上传 " + ext + "格式的文件")
	}

	if err := ginfile.SaveFile(f, fullPath, fileName); err != nil {
		return "", err
	}

	return config.AppConfig.URL + "/" + fullPath + "/" + fileName, nil
}
