package controllers

import (
	"gin_bbs/config"
	"gin_bbs/pkg/ginutils/csrf"
	ginfile "gin_bbs/pkg/ginutils/file"
	"gin_bbs/pkg/ginutils/flash"
	"gin_bbs/pkg/ginutils/oldvalue"
	"gin_bbs/pkg/ginutils/router"
	"gin_bbs/pkg/ginutils/validate"
	"net/http"
	"strconv"

	"gin_bbs/app/auth"
	"gin_bbs/app/viewmodels"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

type (
	renderObj = map[string]interface{}
)

const (
	csrfInputHTML = "csrfInput"
	csrfMetaHTML  = "csrfMeta"
	csrfTokenName = "csrfToken"
)

// Render : 渲染 html
func Render(c *gin.Context, tplPath string, data renderObj) {
	obj := make(pongo2.Context)
	flashStore := flash.Read(c)
	oldValueStore := oldvalue.ReadOldFormValue(c)

	// flash 数据 (使用 flash 后，应该用 redirect，如果用 render 的话得刷新页面后才会出现 flash 信息)
	// 因为 flash 使用 session 实现的，需先写入，再读取，所以完整过程依赖两个 request (oldValue、validate 同理)
	obj[flash.FlashInContextAndCookieKeyName] = flashStore.Data
	// 上次 post form 的数据，用于回填
	obj[oldvalue.OldValueInContextAndCookieKeyName] = oldValueStore.Data
	// 上次表单的验证信息
	errArr, errMap := validate.ReadValidateMessage(c)
	obj[validate.ValidateMessageArrKeyName] = errArr
	obj[validate.ValidateMessageMapKeyName] = errMap
	// csrf
	if config.AppConfig.EnableCsrf {
		if csrfHTML, csrfToken, ok := csrf.CsrfInput(c); ok {
			obj[csrfInputHTML] = csrfHTML
			obj[csrfMetaHTML], _, _ = csrf.CsrfMeta(c)
			obj[csrfTokenName] = csrfToken
		}
	}
	// route class
	obj["route_path"] = c.Request.URL.Path
	// 获取当前登录的用户 (如果用户登录了的话，中间件中会通过 session 存储用户数据)
	if user, err := auth.GetCurrentUserFromContext(c); err == nil {
		obj[config.AppConfig.ContextCurrentUserDataKey] = viewmodels.NewUserViewModelSerializer(user)
	}

	// 填充传递进来的数据
	for k, v := range data {
		obj[k] = v
	}

	c.HTML(http.StatusOK, tplPath+".html", obj)
}

// RenderError : 渲染错误页面
func RenderError(c *gin.Context, code int, title, msg string) {
	errorCode := code
	if code == 419 || code == 403 || code == 429 {
		errorCode = 403
	}

	c.HTML(code, "errors/error.html", pongo2.Context{
		"errorTitle": title,
		"errorMsg":   msg,
		"errorCode":  errorCode,
		"errorImg":   ginfile.StaticPath("/svg/" + strconv.Itoa(errorCode) + ".svg"),
		"backUrl":    router.G("root"),
	})
}

// Render403 -
func Render403(c *gin.Context, msg string) {
	RenderError(c, http.StatusForbidden, msg, msg)
}

// Render404 -
func Render404(c *gin.Context) {
	msg := "很抱歉！您浏览的页面不存在。"
	RenderError(c, http.StatusNotFound, msg, msg)
}

// RenderUnauthorized -
func RenderUnauthorized(c *gin.Context) {
	Render403(c, "很抱歉，您没有权限访问该页面")
}

// RenderTooManyRequests -
func RenderTooManyRequests(c *gin.Context) {
	RenderError(c, 429, "太多请求", "很抱歉！您向我们的服务器发出太多请求了。")
}
