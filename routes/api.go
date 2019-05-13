package routes

import (
	"gin_bbs/pkg/ginutils/router"

	vericode "gin_bbs/app/controllers/api/verification_code"

	"github.com/gin-gonic/gin"
)

func registerAPI(r *router.MyRoute, middlewares ...gin.HandlerFunc) {
	r = r.Group(ApiRoot, middlewares...)

	r.Register("POST", "api.verificationCodes.store", "verificationCodes", vericode.Store)
}
