package vericode

import (
	"gin_bbs/app/cache"
	"gin_bbs/app/controllers"
	"gin_bbs/app/helpers"
	"gin_bbs/app/requests"
	"gin_bbs/pkg/constants"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils/utils"
	"gin_bbs/pkg/ginutils/validate"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Store 手机号码注册用户
func Store(c *gin.Context) {
	var p struct {
		Phone string `json:"phone"`
	}
	if err := c.BindJSON(&p); err != nil {
		controllers.SendErrorResponse(c, errno.ParamsError)
		return
	}

	ok, _, errMap := validate.RunSingle("phone",
		[]validate.ValidatorFunc{
			validate.RequiredValidator(p.Phone),
			validate.PhoneValidator(p.Phone),
			requests.PhoneUniqueValidator(p.Phone),
		},
		[]string{"手机号不能为空", "手机号格式错误"})

	if !ok {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, errMap))
		return
	}

	// 生成 4 位随机数，左侧补 0
	code := strconv.Itoa(utils.RandInt(1, 9999))
	code, _ = utils.LeftPad(code, 4, '0')
	result := helpers.SendSms(p.Phone, code) // 发送短信
	if result.Code == -1 {
		controllers.SendErrorResponse(c, errno.New(errno.SmsError, result))
		return
	}

	// 存入缓存 (十分钟过期)
	expiredAt := 10 * time.Minute
	key := "verificationCode_" + string(utils.RandomCreateBytes(15))
	cache.Put(key, map[string]interface{}{"phone": p.Phone, "code": code}, expiredAt)

	controllers.SendOKResponse(c, map[string]string{
		"key":        key,
		"expired_at": time.Now().Add(expiredAt).Format(constants.DateTimeLayout),
	})
}
