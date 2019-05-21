package link

import (
	. "gin_bbs/app/controllers"
	linkModel "gin_bbs/app/models/link"
	"gin_bbs/app/viewmodels"

	"github.com/gin-gonic/gin"
)

// Index 资源链接列表
func Index(c *gin.Context) {
	all, _ := linkModel.All()
	SendOKResponse(c, ListData{
		List: viewmodels.LinkAPIList(all),
	})
}
