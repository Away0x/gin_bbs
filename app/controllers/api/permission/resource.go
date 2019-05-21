package permission

import (
	. "gin_bbs/app/controllers"
	permissionModel "gin_bbs/app/models/permission"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/app/viewmodels"

	"github.com/gin-gonic/gin"
)

// Index 用户权限列表
func Index(c *gin.Context, currentUser *userModel.User, tokenString string) {
	all, _ := permissionModel.GetUserAllPermission(currentUser)
	SendOKResponse(c, ListData{
		List: viewmodels.PermissionAPIList(all),
	})
}
