package ginutils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Redirect 重定向
func Redirect(c *gin.Context, redirectPath string) {
	// 千万注意，这个地方不能用 301(永久重定向)
	c.Redirect(http.StatusFound, redirectPath)
}
