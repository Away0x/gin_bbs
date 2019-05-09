package routes

import (
	"gin_bbs/pkg/ginutils/router"

	"github.com/gin-gonic/gin"
)

func registerAPI(r *router.MyRoute, middlewares ...gin.HandlerFunc) {
	r = r.Group(ApiRoot, middlewares...)

	r.Register("GET", "api.test", "test", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "api test"})
	})
}
