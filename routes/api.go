package routes

import (
	"gin_bbs/pkg/ginutils/router"
	"gin_bbs/routes/middleware"
	"time"

	vericode "gin_bbs/app/controllers/api/verification_code"

	"github.com/gin-gonic/gin"
)

func registerAPI(r *router.MyRoute, middlewares ...gin.HandlerFunc) {
	r = r.Group(ApiRoot, middlewares...)

	r.Register("POST", "api.verificationCodes.store", "verificationCodes",
		middleware.RateLimiter(1*time.Minute, 1), // 1 分钟 1 次
		vericode.Store)
}
