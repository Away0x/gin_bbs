package last

import (
	"time"

	"github.com/gin-gonic/gin"
)

// LastMiddleware 记录上一次访问的请求信息
func LastMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		k := c.ClientIP()
		data := &LastRequestData{
			Path: c.Request.URL.Path,
		}

		last, ok := lastReqCache.Get(k)
		if ok && last != nil {
			if c.Keys == nil {
				c.Keys = make(map[string]interface{})
			}

			c.Keys[ContextKeysName] = last
		}
		lastReqCache.Set(k, data, time.Hour)

		c.Next()
	}
}
