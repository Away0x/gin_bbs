package routes

import (
	"gin_bbs/app/controllers"
	"gin_bbs/pkg/ginutils/router"

	"fmt"
	ginfile "gin_bbs/pkg/ginutils/file"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RenderAdminIndexPage 渲染管理员后台 index.html
func RenderAdminIndexPage(c *gin.Context) {
	c.Status(http.StatusOK)

	tplPath := ginfile.PublicPath("admin-index.html")
	htmlStr, err := ginfile.ReadFile(tplPath)
	if err != nil {
		fmt.Println(err)
		controllers.Render404(c)
		return
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(c.Writer, htmlStr)
}

func registerAdmin(r *router.MyRoute, middlewares ...gin.HandlerFunc) {
	// 管理员后台静态文件
	rweb := r.Group(AdminWebRoot, middlewares...)
	rweb.Router.Static("", "admin-fe/dist") // 如果想 vue 使用 history 模式，需自己实现一个静态文件服务，没有文件就渲染 index.html
	// 可参考 gin.Context 的 File 方法 (内部是 http.ServeFile)

	// 管理员后台 api
	rapi := r.Group(AdminApiRoot, middlewares...)
	rapi.Register("GET", "admin.api.test", "test", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "admin api test"})
	})
}
