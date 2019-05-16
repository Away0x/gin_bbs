package routes

import (
	"gin_bbs/pkg/ginutils/router"
	"gin_bbs/routes/middleware"
	"time"

	"gin_bbs/app/controllers/api/authorization"
	"gin_bbs/app/controllers/api/captcha"
	"gin_bbs/app/controllers/api/user"
	vericode "gin_bbs/app/controllers/api/verification_code"

	"github.com/gin-gonic/gin"
)

func registerAPI(r *router.MyRoute, middlewares ...gin.HandlerFunc) {
	r = r.Group(APIRoot, middlewares...)

	// 短信验证码
	r.Register("POST", "api.verificationCodes.store", "/verificationCodes",
		middleware.RateLimiter(1*time.Minute, 10), // 1 分钟 10 次
		vericode.Store)
	// 用户注册
	r.Register("POST", "api.users.store", "/users",
		middleware.RateLimiter(1*time.Minute, 10), // 1 分钟 10 次
		user.Store)
	// 图片验证码
	r.Register("POST", "api.captchas.store", "/captchas", captcha.Store)

	// 第三方登录
	r.Register("POST", "api.socials.authorizations.store", "/socials/authorizations/:social_type", authorization.SocialStore)
	// 登录 签发 token
	r.Register("POST", "api.authorizations.store", "/authorizations", authorization.Store)
	// 刷新 token
	r.Register("PUT", "api.authorizations.update", "/authorizations/current", authorization.Update)
	// 刷新 token
	r.Register("DELETE", "api.authorizations.destroy", "/authorizations/current", authorization.Destroy)
}
