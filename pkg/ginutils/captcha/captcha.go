package captcha

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

type Captcha struct {
	ID       string
	ImageURL string
	AudioURL string
}

// New 获取验证码
func New(baseURL string) Captcha {
	id := captcha.New()

	return Captcha{
		ID:       id,
		ImageURL: baseURL + "/" + id + ".png",
		AudioURL: baseURL + "/" + id + ".wav",
	}
}

// Verify 验证
func Verify(captchaID, captchaVal string) bool {
	return captcha.VerifyString(captchaID, captchaVal)
}

// Handler gin handler
func Handler(c *gin.Context) {
	ServeHTTP(c.Writer, c.Request)
}
