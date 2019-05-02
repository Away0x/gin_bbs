package bootstrap

import (
	"gin_bbs/config"
	"gin_bbs/routes"
	"gin_bbs/routes/named"

	"github.com/gin-gonic/gin"
)

func SetupRouter(g *gin.Engine) {
	routes.Register(g)
	printRoute()
}

// 打印命名路由
func printRoute() {
	// 只有非 release 时才可用该函数
	if config.AppConfig.RunMode == config.RunmodeRelease {
		return
	}

	named.PrintRoutes()
}
