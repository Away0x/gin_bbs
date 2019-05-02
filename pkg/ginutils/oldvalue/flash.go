package oldvalue

import (
	"gin_bbs/pkg/ginutils/flash"

	"github.com/gin-gonic/gin"
)

const (
	// OldValueInContextAndCookieKeyName : gin context keys 中的 name (cookie 中也是)
	OldValueInContextAndCookieKeyName = "oldValue"
)

// SaveOldFormValue : 存储上次表单 post 的数据
func SaveOldFormValue(c *gin.Context, obj map[string]string) {
	f := flash.NewFlashByName(OldValueInContextAndCookieKeyName)
	f.Data = obj
	f.SaveByName(c, OldValueInContextAndCookieKeyName)
}

// ReadOldFormValue : 读取上次表单 post 的数据
func ReadOldFormValue(c *gin.Context) *flash.FlashData {
	return flash.ReadByName(c, OldValueInContextAndCookieKeyName)
}
