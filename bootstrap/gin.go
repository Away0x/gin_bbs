package bootstrap

import (
	"gin_bbs/config"

	"gin_bbs/pkg/pongo2gin"

	"github.com/gin-gonic/gin"
)

// SetupGin gin setup
func SetupGin(g *gin.Engine) {
	// 启动模式配置
	gin.SetMode(config.AppConfig.RunMode)

	// 项目静态文件配置
	g.Static("/"+config.AppConfig.PublicPath, config.AppConfig.PublicPath)
	g.StaticFile("/favicon.ico", config.AppConfig.PublicPath+"/favicon.ico")

	// 模板配置
	setupTemplate(g)
}

func setupTemplate(g *gin.Engine) {
	g.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateDir: config.AppConfig.ViewsPath,
		ContentType: "text/html; charset=utf-8",
	})

	// // 注册模板函数
	// g.SetFuncMap(template.FuncMap{
	// 	// 根据 laravel-mix 的 public/mix-manifest.json 生成静态文件 path
	// 	"Mix": helpers.Mix,
	// 	// 生成项目静态文件地址
	// 	"Static": helpers.Static,
	// 	// 获取命名路由的 path
	// 	"Route":         named.G,
	// 	"RelativeRoute": named.GR,
	// })
	// // 模板存储路径
	// g.LoadHTMLGlob(config.AppConfig.ViewsPath + "/**/*")
}
