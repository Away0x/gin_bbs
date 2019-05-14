package vericode

import (
	"gin_bbs/app/cache"
	"gin_bbs/app/controllers"
	"gin_bbs/app/helpers"
	"gin_bbs/app/requests"
	"gin_bbs/pkg/constants"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Store 发送短信
func Store(c *gin.Context) {
	phone, ok := requests.RunPhoneValidate(c)
	if !ok {
		return
	}

	// 生成 4 位随机数，左侧补 0
	code := strconv.Itoa(utils.RandInt(1, 9999))
	code, _ = utils.LeftPad(code, 4, '0')
	result := helpers.SendSms(phone, code) // 发送短信
	if result.Code == -1 {
		controllers.SendErrorResponse(c, errno.New(errno.SmsError, result))
		return
	}

	// 存入缓存 (十分钟过期)
	expiredAt := 10 * time.Minute
	key := "verificationCode_" + string(utils.RandomCreateBytes(15))
	cache.PutStringMap(key, map[string]string{"phone": phone, "code": code}, expiredAt)

	controllers.SendOKResponse(c, map[string]interface{}{
		"key":              key,
		"expired_at":       time.Now().Add(expiredAt).Format(constants.DateTimeLayout),
		"debug_sms_result": result,
	})
}
