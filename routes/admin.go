package routes

import (
	"gin_bbs/app/controllers"
	"gin_bbs/pkg/ginutils/router"

	"fmt"
	ginfile "gin_bbs/pkg/ginutils/file"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	adminFeDistFilePath = "admin-fe/dist"
)

// RenderAdminIndexPage 渲染管理员后台 index.html
func RenderAdminIndexPage(c *gin.Context) {
	c.Status(http.StatusOK)

	htmlStr, err := ginfile.ReadFile(adminFeDistFilePath + "/index.html")
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
	// rweb.Router.Static("", "admin-fe/dist")
	rweb.Register("GET", "admin.index", "/*filepath", func(c *gin.Context) {
		filepath := c.Param("filepath")
		if filepath == "" || filepath == "/" {
			RenderAdminIndexPage(c)
			return
		}

		allPath := adminFeDistFilePath + filepath
		if ginfile.IsExist(allPath) {
			c.File(allPath)
			return
		}

		RenderAdminIndexPage(c)
		return
	})
	// 可参考 gin.Context 的 File 方法 (内部是 http.ServeFile)

	// 管理员后台 api
	rapi := r.Group(AdminApiRoot, middlewares...)
	rapi.Register("GET", "admin.api.test", "test", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "admin api test"})
	})
}
