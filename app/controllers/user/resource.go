package user

import (
	"gin_bbs/app/controllers"
	"gin_bbs/pkg/ginutils"

	userModel "gin_bbs/app/models/user"
	"gin_bbs/app/viewmodels"

	"github.com/gin-gonic/gin"
)

// 展示用户信息页面
func Show(c *gin.Context, currentUser *userModel.User) {
	id, err := ginutils.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	// 如果要看的就是当前用户，那么就不用再去数据库中获取了
	user := currentUser
	if id != int(currentUser.ID) {
		user, err = userModel.Get(id)
	}

	if err != nil || user == nil {
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
