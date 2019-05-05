package helpers

import (
	"errors"
	"gin_bbs/config"
	ginfile "gin_bbs/pkg/ginutils/file"
	"mime/multipart"
	"path"
	"strings"
)

const (
	// UploadFolder 上传文件的文件夹路径
	UploadFolder = "/uploads/"
	// ImagesUploadFolder 图片存储路径
	ImagesUploadFolder = UploadFolder + "images/"
)

// SaveImage 保存图片
// filePrefix: 文件名前缀
// maxWidth: 图片最大宽度，0 为不限制
func SaveImage(f *multipart.FileHeader, folderName, filePrefix string, maxWidth int) (string, error) {
	fileName, ext := ginfile.CreateRandomFileName(f, filePrefix, ".png")
	folderPath := ginfile.PublicPath(ImagesUploadFolder + folderName + ginfile.CreateBaseTimeFolderName())

	ext = strings.ToLower(ext)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" && ext != ".bmp" && ext != ".gif" {
		return "", errors.New("文件格式错误，不能上传 " + ext + "格式的文件")
	}

	// 保存
	src, err := f.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	if err := ginfile.SaveFile(src, folderPath, fileName); err != nil {
		return "", err
	}

	filePath := path.Join(folderPath, fileName)
	// 需要 resize 图像
	if maxWidth > 0 {
		ginfile.ReduceImageSize(filePath, maxWidth)
	}

	return config.AppConfig.URL + "/" + filePath, nil
}
