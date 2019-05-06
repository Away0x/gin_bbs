package routes

import (
	"gin_bbs/pkg/ginutils/csrf"
	"gin_bbs/pkg/ginutils/last"
	"gin_bbs/pkg/ginutils/oldvalue"
	"gin_bbs/pkg/ginutils/session"

	"gin_bbs/pkg/ginutils/router"
	"gin_bbs/routes/middleware"

	"gin_bbs/app/controllers"

	"github.com/gin-gonic/gin"
)

// Register 注册路由和中间件
func Register(g *gin.Engine) *gin.Engine {
	// ---------------------------------- 注册全局中间件 ----------------------------------
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(last.LastMiddleware()) // 记录上一次请求信息

	// ---------------------------------- 注册路由 ----------------------------------
	r := &router.MyRoute{Router: g}

	// +++++++++++++++++++ error +++++++++++++++++++
	g.NoRoute(func(c *gin.Context) {
		controllers.Render404(c)
	})
	g.NoMethod(func(c *gin.Context) {
		controllers.Render404(c)
	})

	// +++++++++++++++++++ web +++++++++++++++++++
	registerWeb(r,
		// session (flash 功能也依赖该 middleware)
		session.SessionMiddleware(),
		// csrf
		csrf.Middleware(func(c *gin.Context, _ bool) {
			if c.GetHeader("X-Requested-With") == "XMLHttpRequest" {
				c.JSON(403, gin.H{"msg": "很抱歉！您的 Session 已过期，请刷新后再试一次。"})
			} else {
				controllers.Render403(c, "很抱歉！您的 Session 已过期，请刷新后再试一次。")
			}
		}),
		// 记忆上次表单提交的内容，消费即消失
		oldvalue.OldValueMiddleware(),
		// 中间件中会从 session 中获取到 current user model
		middleware.CurrentUserMiddleware(),
	)

	// +++++++++++++++++++ api +++++++++++++++++++
	registerAPI(r)

	return g
}
