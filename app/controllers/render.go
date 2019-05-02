package controllers

import (
	"gin_bbs/app/helpers"
	"gin_bbs/config"
	"gin_bbs/pkg/flash"
	"gin_bbs/routes/named"
	"net/http"
	"strconv"

	"gin_bbs/pkg/ginutils"

	"github.com/gin-gonic/gin"
)

type (
	renderObj = map[string]interface{}
)

const (
	csrfInputHTML = "csrfField"
	csrfTokenName = "csrfToken"
)

// Render : 渲染 html
func Render(c *gin.Context, tplPath string, data renderObj) {
	obj := make(renderObj)
	flashStore := flash.Read(c)
	oldValueStore := flash.ReadOldFormValue(c)
	validateMsgArr := flash.ReadValidateMessage(c)

	// flash 数据
	obj[flash.FlashInContextAndCookieKeyName] = flashStore.Data
	// 上次 post form 的数据，用于回填
	obj[flash.OldValueInContextAndCookieKeyName] = oldValueStore.Data
	// 上次表单的验证信息
	obj[flash.ValidateContextAndCookieKeyName] = validateMsgArr
	// csrf
	if config.AppConfig.EnableCsrf {
		if csrfHTML, csrfToken, ok := ginutils.CsrfField(c, config.AppConfig.CsrfParamName); ok {
			obj[csrfInputHTML] = csrfHTML
			obj[csrfTokenName] = csrfToken
		}
	}

	// 填充传递进来的数据
	for k, v := range data {
		obj[k] = v
	}

	c.HTML(http.StatusOK, tplPath, obj)
}

// RenderError : 渲染错误页面
func RenderError(c *gin.Context, code int, msg string) {
	errorCode := code
	if code == 419 || code == 403 {
		errorCode = 403
	}

	c.HTML(code, "error/error.html", gin.H{
		"errorMsg":  msg,
		"errorCode": errorCode,
		"errorImg":  helpers.Static("/svg/" + strconv.Itoa(code) + ".svg"),
		"backUrl":   named.G("root"),
	})
}

// Render403 -
func Render403(c *gin.Context, msg string) {
	RenderError(c, http.StatusForbidden, msg)
}

// Render404 -
func Render404(c *gin.Context) {
	RenderError(c, http.StatusNotFound, "很抱歉！您浏览的页面不存在。")
}

// RenderUnauthorized -
func RenderUnauthorized(c *gin.Context) {
	Render403(c, "很抱歉，您没有权限访问该页面")
}
