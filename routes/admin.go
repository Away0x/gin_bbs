package routes

import (
	"gin_bbs/pkg/ginutils/router"

	"gin_bbs/app/controllers/admin"

	"github.com/gin-gonic/gin"
)

func registerAdmin(r *router.MyRoute, middlewares ...gin.HandlerFunc) {
	r = r.Group("/admin", middlewares...)

	r.Register("GET", "admin.index", "", admin.Index)
}
