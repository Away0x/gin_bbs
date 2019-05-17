package middleware

import (
	"gin_bbs/app/controllers"
	"gin_bbs/pkg/constants"
	"gin_bbs/pkg/errno"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"golang.org/x/time/rate"
)

// RateLimiter 限制指定时间内最多多少个请求，超出会进入错误页面
func RateLimiter(duration time.Duration, n int) gin.HandlerFunc {
	// 默认 5m 过期，每 10m 清除缓冲中的过期 key
	// 定期清除缓存中的过期 key，是通过一个常驻 goroutine 实现的
	limiterCache := cache.New(5*time.Minute, 10*time.Minute)

	return func(c *gin.Context) {
		k := c.ClientIP() // limit rate by client ip

		limiter, ok := limiterCache.Get(k)
		if !ok {
			var expire time.Duration
			// limiter liveness time duration is 1 hour
			// ip 限制 duration 时间内最多 n 个请求
			limiter, expire = rate.NewLimiter(rate.Every(duration), n), time.Hour
			limiterCache.Set(k, limiter, expire)
		}

		ok = limiter.(*rate.Limiter).Allow()
		if !ok {
			if constants.IsApiRequest(c) {
				controllers.SendErrorResponse(c, errno.TooManyRequestError)
			} else {
				controllers.RenderTooManyRequests(c)
			}
			c.AbortWithStatus(429) // handle exceed rate limit request
			return
		}
		c.Next()
	}
}

// MaxAllowed -
// func MaxAllowed(n int) gin.HandlerFunc {
// 	sem := make(chan struct{}, n)
// 	acquire := func() { sem <- struct{}{} }
// 	release := func() { <-sem }

// 	return func(c *gin.Context) {
// 		acquire()       // before request
// 		defer release() // after request
// 		c.Next()
// 	}
// }
