package controllers

import (
	"gin_bbs/config"
	"gin_bbs/pkg/ginutils"
	"gin_bbs/pkg/ginutils/router"

	"github.com/gin-gonic/gin"
)

// Redirect : 路由重定向 use path
func Redirect(c *gin.Context, redirectPath string, withRoot bool) {
	path := redirectPath
	if withRoot {
		path = config.AppConfig.URL + redirectPath
	}

	ginutils.Redirect(c, path)
}

// RedirectRouter : 路由重定向 use router name
func RedirectRouter(c *gin.Context, routerName string, args ...interface{}) {
	ginutils.Redirect(c, router.G(routerName, args...))
}
