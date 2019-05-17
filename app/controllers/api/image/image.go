package image

import (
	userModel "gin_bbs/app/models/user"

	"github.com/gin-gonic/gin"
)

// Store 上传图片
func Store(c *gin.Context, currentUser *userModel.User, tokenString string) {

}
