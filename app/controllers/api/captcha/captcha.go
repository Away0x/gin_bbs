package captcha

import (
	"gin_bbs/app/cache"
	"gin_bbs/app/controllers"
	"gin_bbs/app/requests"
	"gin_bbs/config"
	"gin_bbs/pkg/constants"
	"gin_bbs/pkg/ginutils/captcha"
	"gin_bbs/pkg/ginutils/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// Store 图片验证码
func Store(c *gin.Context) {
	phone, ok := requests.RunPhoneValidate(c)
	if !ok {
		return
	}

	captcha := captcha.New("/captcha")
	expiredAt := 2 * time.Minute
	key := "captcha_" + string(utils.RandomCreateBytes(15))
	cache.PutStringMap(key, map[string]string{"phone": phone, "captcha_id": captcha.ID}, expiredAt)

	controllers.SendOKResponse(c, map[string]interface{}{
		"captcha_key":           key,
		"expired_at":            time.Now().Add(expiredAt).Format(constants.DateTimeLayout),
		"captcha_image_content": config.AppConfig.URL + captcha.ImageURL,
	})
}
