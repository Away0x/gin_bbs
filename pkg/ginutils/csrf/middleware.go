package csrf

import (
	"gin_bbs/pkg/ginutils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware : csrf middleware
// errorFn: csrf 验证不通过时执行的 handler
func Middleware(errorFn func(*gin.Context, bool)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if ginutils.GetGinUtilsConfig().EnableCsrf {
			// cookie 中获取 csrf token (如没有则设置)
			csrfToken := getCsrfTokenFromCookie(c)
			method := c.Request.Method

			// 非 GET 并且开启了 csrf
			if method == http.MethodPost ||
				method == http.MethodDelete ||
				method == http.MethodPut ||
				method == http.MethodPatch {
				// params 中获取 csrf token
				paramCsrfToken, inHeader := getCsrfTokenFromParamsOrHeader(c)

				if paramCsrfToken == "" || paramCsrfToken != csrfToken {
					errorFn(c, inHeader)
					c.Abort()
					return
				}
			}
		}

		c.Next()
	}
}
