package routes

import (
	"gin_bbs/pkg/ginutils/csrf"
	"gin_bbs/pkg/ginutils/last"
	"gin_bbs/pkg/ginutils/oldvalue"
	"gin_bbs/pkg/ginutils/session"

	"gin_bbs/pkg/constants"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils/router"
	"gin_bbs/routes/middleware"

	"gin_bbs/app/controllers"
	"gin_bbs/config"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const (
	// APIRoot -
	APIRoot = "/api"

	// AdminWebRoot -
	AdminWebRoot = "/admin"
	// AdminAPIRoot -
	AdminAPIRoot = "/admin-api"
)

// Register 注册路由和中间件
func Register(g *gin.Engine) *gin.Engine {
	// ---------------------------------- 注册全局中间件 ----------------------------------
	g.Use(gin.Recovery())
	if config.AppConfig.RunMode != config.RunmodeRelease {
		g.Use(gin.Logger())
	}
	g.Use(last.LastMiddleware()) // 记录上一次请求信息

	// ---------------------------------- 注册路由 ----------------------------------
	r := &router.MyRoute{Router: g}

	// +++++++++++++++++++ swagger +++++++++++++++++++
	// 需全局安装 go get -u github.com/swaggo/swag/cmd/swag 然后 swag init 生成文档
	if config.AppConfig.RunMode != config.RunmodeRelease {
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// +++++++++++++++++++ web +++++++++++++++++++
	registerWeb(r,
		// session
		session.SessionMiddleware(),
		// csrf
		csrf.Middleware(func(c *gin.Context, _ bool) {
			if c.GetHeader(constants.HeaderRequestedWith) != "" {
				controllers.SendErrorResponse(c, errno.SessionError)
			} else {
				controllers.Render403(c, "很抱歉！您的 Session 已过期，请刷新后再试一次。")
			}
		}),
		// 记忆上次表单提交的内容，消费即消失
		oldvalue.OldValueMiddleware(),
		// 中间件中会从 session 中获取到 current user model
		middleware.CurrentUserMiddleware(),
	)

	// +++++++++++++++++++ admin +++++++++++++++++++
	registerAdmin(r)

	// +++++++++++++++++++ api +++++++++++++++++++
	registerAPI(r)

	// ---------------------------------- error ----------------------------------
	g.NoRoute(func(c *gin.Context) {
		if c.GetHeader(constants.HeaderRequestedWith) != "" {
			controllers.SendErrorResponse(c, errno.NotFoundError)
		} else {
			controllers.Render404(c)
		}
	})
	g.NoMethod(func(c *gin.Context) {
		if c.GetHeader(constants.HeaderRequestedWith) != "" {
			controllers.SendErrorResponse(c, errno.NotFoundError)
		} else {
			controllers.Render404(c)
		}
	})

	return g
}
