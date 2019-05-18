package image

import (
	"gin_bbs/app/controllers"
	userModel "gin_bbs/app/models/user"
	request "gin_bbs/app/requests/api/image"
	"gin_bbs/app/viewmodels"

	"github.com/gin-gonic/gin"
)

// Store 上传图片
func Store(c *gin.Context, currentUser *userModel.User, tokenString string) {
	img, _ := c.FormFile("image")
	req := &request.Upload{
		Type:  c.PostForm("type"),
		Image: img,
	}

	image, err := req.Run(currentUser)
	if err != nil {
		controllers.SendErrorResponse(c, err)
		return
	}

	controllers.SendOKResponse(c, viewmodels.Image(image))
}
