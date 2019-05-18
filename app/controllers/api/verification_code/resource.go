package vericode

import (
	"gin_bbs/app/cache"
	"gin_bbs/app/controllers"
	"gin_bbs/app/helpers"
	vericode "gin_bbs/app/requests/api/verification_code"
	"gin_bbs/pkg/constants"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Store 发送短信
// @Summary 发送短信
// @Tags verificationCodes
// @Accept  json
// @Produce  json
// @Param req body vericode.VerificationCode true "req"
// @Success 200 {object} controllers.Response "{"key": "verificationCode_xxxxx","debug_sms_result": "xxxx","expired_at": "2019-05-15 17:23:21"}"
// @Router /api/verificationCodes [post]
func Store(c *gin.Context) {
	// 验证图片验证码
	var req vericode.VerificationCode
	if err := c.ShouldBind(&req); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}
	phone, err := req.Run()
	if err != nil {
		controllers.SendErrorResponse(c, err)
		return
	}

	// 发送短信
	code := strconv.Itoa(utils.RandInt(1, 9999)) // 生成 4 位随机数，左侧补 0
	code, _ = utils.LeftPad(code, 4, '0')
	result := helpers.SendSms(phone, code)
	if result.Code == -1 {
		controllers.SendErrorResponse(c, errno.New(errno.SmsError, result))
		return
	}

	// 短信 code 存入缓存 (十分钟过期)
	expiredAt := 10 * time.Minute
	key := "verificationCode_" + string(utils.RandomCreateBytes(15))
	cache.PutStringMap(key, map[string]string{"phone": phone, "code": code}, expiredAt)

	controllers.SendOKResponse(c, map[string]interface{}{
		"key":              key,
		"expired_at":       time.Now().Add(expiredAt).Format(constants.DateTimeLayout),
		"debug_sms_result": result,
	})
}
