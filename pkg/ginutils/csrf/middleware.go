package csrf

import (
	"gin_bbs/pkg/ginutils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CsrfMiddleware : csrf middleware
func CsrfMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if ginutils.GetGinUtilsConfig().EnableCsrf {
			// cookie 中获取 csrf token (如没有则设置)
			csrfToken := getCsrfTokenFromCookie(c)

			// POST 并且开启了 csrf
			if c.Request.Method == http.MethodPost {
				// params 中获取 csrf token
				paramCsrfToken := getCsrfTokenFromParamsOrHeader(c)

				if paramCsrfToken == "" || paramCsrfToken != csrfToken {
					ginutils.GetGinUtilsConfig().CsrfErrorHandler(c)
					c.Abort()
					return
				}
			}
		}

		c.Next()
	}
}
