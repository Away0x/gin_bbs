package routes

import (
	"gin_bbs/pkg/ginutils/router"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func registerWeb(r *router.MyRoute) {
	r.Register("GET", "root", "/", func(c *gin.Context) {
		c.HTML(200, "index.html", pongo2.Context{})
	})
	r.Register("GET", "index", "/index", func(c *gin.Context) {
		c.HTML(200, "index.html", pongo2.Context{})
	})
	r.Register("POST", "index-post", "/index", func(c *gin.Context) {
		c.HTML(200, "index.html", pongo2.Context{})
	})

	test := r.Group("test")
	{
		test.Register("GET", "test-1", "/one", func(c *gin.Context) {
			c.String(200, "test1")
		})
		test.Register("GET", "test-2", "/two", func(c *gin.Context) {
			c.String(200, "test2")
		})
	}
}
