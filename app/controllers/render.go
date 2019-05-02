package controllers

import (
	"gin_bbs/config"
	"gin_bbs/pkg/ginutils/csrf"
	"gin_bbs/pkg/ginutils/flash"
	"gin_bbs/pkg/ginutils/oldvalue"
	"gin_bbs/pkg/ginutils/validate"
	"net/http"

	"gin_bbs/app/auth"
	viewmodels "gin_bbs/app/view_models"

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
	// 想要获取消息闪现的话，得用 redirect，不能重新 render
	// 因为这里会消费掉上次的 flash
	flashStore := flash.Read(c)
	oldValueStore := oldvalue.ReadOldFormValue(c)
	validateMsgArr := validate.ReadValidateMessage(c)

	// flash 数据
	obj[flash.FlashInContextAndCookieKeyName] = flashStore.Data
	// 上次 post form 的数据，用于回填
	obj[oldvalue.OldValueInContextAndCookieKeyName] = oldValueStore.Data
	// 上次表单的验证信息
	obj[validate.ValidateContextAndCookieKeyName] = validateMsgArr
	// csrf
	if config.AppConfig.EnableCsrf {
		if csrfHTML, csrfToken, ok := csrf.CsrfInput(c); ok {
			obj[csrfInputHTML] = csrfHTML
			obj[csrfMetaHTML], _, _ = csrf.CsrfMeta(c)
			obj[csrfTokenName] = csrfToken
		}
	}
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
