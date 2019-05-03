package last

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

const (
	ContextKeysName = "lastTimeRequest"
)

var lastReqCache = cache.New(5*time.Minute, 10*time.Minute)

type LastRequestData struct {
	Path string
}

func Read(c *gin.Context) *LastRequestData {
	l := c.Keys[ContextKeysName]
	if l == nil {
		return nil
	}

	if result, ok := l.(*LastRequestData); ok {
		return result
	}

	return nil
}
