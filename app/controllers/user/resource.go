package user

import (
	"gin_bbs/app/controllers"

	"gin_bbs/app/auth"
	userModel "gin_bbs/app/models/user"
	userRequest "gin_bbs/app/requests/user"
	"gin_bbs/app/viewmodels"
	"gin_bbs/pkg/ginutils"
	"gin_bbs/pkg/ginutils/flash"

	"gin_bbs/app/policies"

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
func Edit(c *gin.Context, currentUser *userModel.User) {
	id, err := ginutils.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	// 只能更新自己
	if ok := policies.UserPolicyUpdate(c, currentUser, id); !ok {
		return
	}

	controllers.Render(c, "users/edit", gin.H{})
}

// 更新用户
func Update(c *gin.Context, currentUser *userModel.User) {
	id, err := ginutils.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	// 只能更新自己
	if ok := policies.UserPolicyUpdate(c, currentUser, id); !ok {
		return
	}

	// 验证参数并更新用户
	file, _ := c.FormFile("avatar")
	userForm := &userRequest.UserUpdateForm{
		ID:           id,
		Name:         c.PostForm("name"),
		Email:        c.PostForm("email"),
		Introduction: c.PostForm("introduction"),
		Avatar:       file,
	}

	if ok := userForm.ValidateAndUpdate(c, currentUser); !ok {
		controllers.RedirectRouter(c, "users.edit", currentUser.ID)
		return
	}

	flash.NewSuccessFlash(c, "个人资料更新成功！")
	controllers.RedirectRouter(c, "users.show", currentUser.ID)
}
