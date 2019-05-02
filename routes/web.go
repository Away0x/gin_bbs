package routes

import (
	"gin_bbs/pkg/ginutils/router"

	"gin_bbs/app/controllers/auth/login"
	"gin_bbs/app/controllers/auth/password"
	"gin_bbs/app/controllers/auth/register"
	"gin_bbs/app/controllers/auth/verification"
	"gin_bbs/app/controllers/page"
)

func registerWeb(r *router.MyRoute) {
	r.Register("GET", "root", "/", page.Root)

	// ------------------------------------- Auth -------------------------------------
	// 用户身份验证相关的路由
	r.Register("GET", "login.show", "/login", login.ShowLoginForm)
	r.Register("POST", "login", "/login", login.Login)
	r.Register("POST", "logout", "/logout", login.Logout)

	// 用户注册相关路由
	r.Register("GET", "register.show", "/register", register.ShowRegistrationForm)
	r.Register("POST", "register", "/register", register.Register)

	// 密码重置相关路由
	pwdRouter := r.Group("/password")
	{
		pwdRouter.Register("GET", "password.request", "/reset", password.ShowLinkRequestForm)
		pwdRouter.Register("POST", "password.email", "/email", password.SendResetLinkEmail)
		pwdRouter.Register("GET", "password.reset", "/reset/:token", password.ShowResetForm)
		pwdRouter.Register("POST", "password.update", "/reset", password.Reset)
	}

	// Email 认证相关路由
	verificationRouter := r.Group("/email")
	{
		verificationRouter.Register("GET", "verification.notice", "/verify", verification.Show)
		verificationRouter.Register("GET", "verification.verify", "/verify/:id", verification.Verify)
		verificationRouter.Register("GET", "verification.resend", "/resend", verification.Resend)
	}
}
