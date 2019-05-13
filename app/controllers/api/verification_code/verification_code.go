package vericode

import (
	"gin_bbs/app/controllers"

	"github.com/gin-gonic/gin"
)

func Store(c *gin.Context) {
	controllers.SendOKResponse(c, map[string]string{
		"test_message": "store verification code",
	})
}
