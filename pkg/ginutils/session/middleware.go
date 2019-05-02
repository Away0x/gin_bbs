package session

import (
	"github.com/gin-gonic/gin"
	ginSessions "github.com/tommy351/gin-sessions"
)

var (
	sessionKeyPairs  = []byte("secret123")
	sessionStoreName = "my_session"
)

// SessionMiddleware -
func SessionMiddleware() gin.HandlerFunc {
	store := ginSessions.NewCookieStore(sessionKeyPairs)
	store.Options(ginSessions.Options{
		HttpOnly: true,
		Path:     "/",
		MaxAge:   86400 * 30,
	})

	return ginSessions.Middleware(sessionStoreName, store)
}
