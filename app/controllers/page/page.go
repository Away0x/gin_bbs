package page

import (
	"gin_bbs/app/controllers"

	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
	controllers.Render(c, "pages/root", gin.H{})
}
