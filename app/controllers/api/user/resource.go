package user

import (
	"gin_bbs/app/controllers"
	permissionModel "gin_bbs/app/models/permission"
	userModel "gin_bbs/app/models/user"
	request "gin_bbs/app/requests/api/user"
	"gin_bbs/app/viewmodels"
	"gin_bbs/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Store 用户注册
// @Summary 用户注册
// @Tags users
// @Accept  json
// @Produce  json
// @Param req body user.User true "req"
// @Success 200 {object} controllers.Response "{}"
// @Router /api/users [post]
func Store(c *gin.Context) {
	var req request.Register
	if err := c.ShouldBindJSON(&req); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}

	if err := req.ValidateAndCreateUser(); err != nil {
		controllers.SendErrorResponse(c, err)
		return
	}

	controllers.SendOKResponse(c, nil)
}

// Show 获取用户信息
func Show(c *gin.Context, currentUser *userModel.User, tokenString string) {
	roles, _ := permissionModel.UserRoles(currentUser)
	controllers.SendOKResponse(c, viewmodels.NewUserAPIHasRoles(currentUser, roles))
}

// Update 编辑用户信息
func Update(c *gin.Context, currentUser *userModel.User, tokenString string) {
	var req request.Edit
	if err := c.ShouldBind(&req); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}
	req.UserID = currentUser.ID

	if err := req.Run(currentUser); err != nil {
		controllers.SendErrorResponse(c, err)
		return
	}

	controllers.SendOKResponse(c, viewmodels.NewUserAPISerializer(currentUser))
}
