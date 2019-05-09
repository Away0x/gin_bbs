package topic

import (
	"net/http"

	"gin_bbs/app/helpers"
	userModel "gin_bbs/app/models/user"

	"github.com/gin-gonic/gin"
)

// UploadImage 上传图片
func UploadImage(c *gin.Context, currentUser *userModel.User) {
	data := gin.H{
		"success":   false,
		"msg":       "上传失败",
		"file_path": "",
	}

	// 是否有上传文件
	file, _ := c.FormFile("upload_file")
	if file != nil {
		path, err := helpers.SaveImage(file, "topics", currentUser.GetIDstring(), 1024)
		if err == nil {
			data = gin.H{
				"success":   true,
				"msg":       "上传成功",
				"file_path": path,
			}
		}
	}

	c.JSON(http.StatusOK, data)
}
