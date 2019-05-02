package routes

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func registerWeb(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", pongo2.Context{})
	})
}
