package routes

import (
	"gin_bbs/pkg/ginutils/router"
	"gin_bbs/routes/middleware"
	"gin_bbs/routes/wrapper"
	"time"

	"gin_bbs/app/controllers/api/authorization"
	"gin_bbs/app/controllers/api/captcha"
	"gin_bbs/app/controllers/api/category"
	"gin_bbs/app/controllers/api/image"
	"gin_bbs/app/controllers/api/link"
	"gin_bbs/app/controllers/api/notification"
	"gin_bbs/app/controllers/api/permission"
	"gin_bbs/app/controllers/api/reply"
	"gin_bbs/app/controllers/api/topic"
	"gin_bbs/app/controllers/api/user"
	vericode "gin_bbs/app/controllers/api/verification_code"

	"github.com/gin-gonic/gin"
)

func registerAPI(r *router.MyRoute, middlewares ...gin.HandlerFunc) {
	r = r.Group(APIRoot, middlewares...)

	// ------------------------------------- Auth -------------------------------------
	// +++++++++++++++ 注册、登录、token 相关 +++++++++++++++
	{
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
		r.Register("POST", "api.socials.authorizations.store", "/socials/authorizations/:social_type",
			authorization.SocialStore)
		// 登录 签发 token
		r.Register("POST", "api.authorizations.store", "/authorizations",
			authorization.Store)

		// 刷新 token
		r.Register("PUT", "api.authorizations.update", "/authorizations/current",
			middleware.TokenAuth(),
			middleware.RateLimiter(1*time.Minute, 3), // 1 分钟 3 次
			wrapper.GetToken(authorization.Update))
		// 删除 token
		r.Register("DELETE", "api.authorizations.destroy", "/authorizations/current",
			middleware.TokenAuth(),
			wrapper.GetToken(authorization.Destroy))
	}

	// 获取活跃用户列表
	r.Register("GET", "api.users.actived", "/users/actived", user.ActivedIndex)
	// +++++++++++++++ 用户相关 +++++++++++++++
	userRouter := r.Group("/user", middleware.TokenAuth())
	{
		// 获取当前登录用户信息
		userRouter.Register("GET", "api.user.show", "", wrapper.GetToken(user.Show))
		// 图片资源
		userRouter.Register("POST", "api.images.store", "images", wrapper.GetToken(image.Store))
		// 编辑用户信息
		userRouter.Register("PATCH", "api.user.update", "", wrapper.GetToken(user.Update))
	}

	// ------------------------------------- category -------------------------------------
	catRouter := r.Group("/categories")
	{
		// 分类列表
		catRouter.Register("GET", "api.categories.index", "", category.Index)
	}

	// ------------------------------------- topic -------------------------------------
	topicRouter := r.Group("/topics")
	{
		// 发布话题
		topicRouter.Register("POST", "api.topics.store", "",
			middleware.TokenAuth(),
			wrapper.GetToken(topic.Store))
		// 修改话题
		topicRouter.Register("PATCH", "api.topics.update", "/:id",
			middleware.TokenAuth(),
			wrapper.GetToken(topic.Update))
		// 删除话题
		topicRouter.Register("DELETE", "api.topics.destroy", "/:id",
			middleware.TokenAuth(),
			wrapper.GetToken(topic.Destroy))
		// 话题列表
		topicRouter.Register("GET", "api.topics.index", "", topic.Index)
		// 用户话题列表
		topicRouter.Register("GET", "api.users.topics.index", "/user/:user_id", topic.UserIndex)
		// 话题详情
		topicRouter.Register("GET", "api.topics.show", "/show/:id", topic.Show)
	}

	// ------------------------------------- reply -------------------------------------
	replyRouter := r.Group("/replies")
	{
		// 发表回复
		replyRouter.Register("POST", "api.replies.store", "",
			middleware.TokenAuth(),
			wrapper.GetToken(reply.Store))
		// 删除回复
		replyRouter.Register("DELETE", "api.replies.destroy", "/:id",
			middleware.TokenAuth(),
			wrapper.GetToken(reply.Destroy))
		// 话题回复列表
		replyRouter.Register("GET", "api.topic.replies", "/topic/:topic_id", reply.TopicReplies)
		// 用户回复列表
		replyRouter.Register("GET", "api.user.replies", "/user/:user_id", reply.UserReplies)
	}

	// ------------------------------------- notification -------------------------------------
	notificationRouter := r.Group("/notifications", middleware.TokenAuth())
	{
		// 用户通知列表
		notificationRouter.Register("GET", "api.user.notifications.index", "",
			middleware.TokenAuth(),
			wrapper.GetToken(notification.Index))
		// 通知统计
		notificationRouter.Register("GET", "api.user.notifications.stats", "/stats",
			middleware.TokenAuth(),
			wrapper.GetToken(notification.Stats))
		// 标记消息通知为已读
		notificationRouter.Register("PATCH", "api.user.notifications.read", "/read",
			middleware.TokenAuth(),
			wrapper.GetToken(notification.Read))
	}

	// ------------------------------------- permission -------------------------------------
	permissionRouter := r.Group("/permissions", middleware.TokenAuth())
	{
		// 当前登录用户权限
		permissionRouter.Register("GET", "api.user.permissions.index", "",
			middleware.TokenAuth(),
			wrapper.GetToken(permission.Index))
	}

	// ------------------------------------- link -------------------------------------
	linkRouter := r.Group("/links")
	{
		// 资源链接列表
		linkRouter.Register("GET", "api.links.index", "", link.Index)
	}
}
