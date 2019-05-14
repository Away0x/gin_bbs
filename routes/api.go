package routes

import (
	"gin_bbs/pkg/ginutils/router"
	"gin_bbs/routes/middleware"
	"time"

	"gin_bbs/app/controllers/api/user"
	vericode "gin_bbs/app/controllers/api/verification_code"

	"github.com/gin-gonic/gin"
)

func registerAPI(r *router.MyRoute, middlewares ...gin.HandlerFunc) {
	r = r.Group(ApiRoot, middlewares...)

	// 短信验证码
	r.Register("POST", "api.verificationCodes.store", "verificationCodes",
		middleware.RateLimiter(1*time.Minute, 10), // 1 分钟 10 次
		vericode.Store)
	// 用户注册
	r.Register("POST", "api.users.store", "users",
		middleware.RateLimiter(1*time.Minute, 10), // 1 分钟 10 次
		user.Store)
}
