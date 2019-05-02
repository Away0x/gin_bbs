package login

import (
	"gin_bbs/app/controllers"

	"github.com/gin-gonic/gin"
)

func ShowLoginForm(c *gin.Context) {
	controllers.Render(c, "auth/login", gin.H{})
}

func Login(c *gin.Context) {
	controllers.RedirectRouter(c, "login.show")
}

func Logout(c *gin.Context) {

}
