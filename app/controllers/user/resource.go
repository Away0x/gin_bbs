package user

import (
	"gin_bbs/app/controllers"

	"gin_bbs/app/auth"
	"gin_bbs/app/viewmodels"

	"github.com/gin-gonic/gin"
)

// 展示用户信息页面
func Show(c *gin.Context) {
	user, err := auth.GetUserFromParamIDOrContext(c)
	if err != nil {
		controllers.Render404(c)
		return
	}

	controllers.Render(c, "users/show", gin.H{
		"user": viewmodels.NewUserViewModelSerializer(user),
	})
}

// 编辑用户信息页面
func Edit(c *gin.Context) {
	controllers.Render(c, "users/edit", gin.H{})
}

// 更新用户
func Update(c *gin.Context) {
}
