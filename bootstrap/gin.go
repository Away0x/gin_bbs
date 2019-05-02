package bootstrap

import (
	"gin_bbs/config"
	"path"

	"gin_bbs/pkg/ginutils"
	pongo2utils "gin_bbs/pkg/ginutils/pongo2"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

// SetupGin gin setup
func SetupGin(g *gin.Engine) {
	// 配置 ginutils
	ginutils.InitGinUtils(ginutils.ConfigOption{
		URL: config.AppConfig.URL,

		PublicPath:  config.AppConfig.PublicPath,
		MixFilePath: path.Join(config.AppConfig.PublicPath, "mix-manifest.json"),

		EnableCsrf:     config.AppConfig.EnableCsrf,
		CsrfParamName:  config.AppConfig.CsrfParamName,
		CsrfHeaderName: config.AppConfig.CsrfHeaderName,
		CsrfErrorHandler: func(c *gin.Context) {
			c.String(403, "很抱歉！您的 Session 已过期，请刷新后再试一次。")
		},
	})

	// 启动模式配置
	gin.SetMode(config.AppConfig.RunMode)

	// 项目静态文件配置
	g.Static("/"+config.AppConfig.PublicPath, config.AppConfig.PublicPath)
	g.StaticFile("/favicon.ico", config.AppConfig.PublicPath+"/favicon.ico")

	// 模板配置
	setupTemplate(g)
}

func setupTemplate(g *gin.Engine) {
	// 使用 pongo2 模板
	g.HTMLRender = pongo2utils.New(pongo2utils.RenderOptions{
		TemplateDir: config.AppConfig.ViewsPath,
		ContentType: "text/html; charset=utf-8",
	})

	// 注册模板全局变量
	pongo2.Globals["flashMessage"] = []string{
		"danger", "warning", "success", "info",
	}

	// 注册模板全局 filter
	// pongo2.RegisterFilter("static", helpers.Static)

	// 注册模板全局 tag
	pongo2.RegisterTag("static", pongo2utils.StaticTag)
	pongo2.RegisterTag("mix", pongo2utils.MixTag)
}
