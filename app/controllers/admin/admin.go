package admin

import (
	"gin_bbs/app/controllers"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	controllers.Render(c, "admin/index", gin.H{})
}
