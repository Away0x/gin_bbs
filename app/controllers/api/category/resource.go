package category

import (
	"gin_bbs/app/controllers"
	categoryModel "gin_bbs/app/models/category"
	"gin_bbs/app/viewmodels"

	"github.com/gin-gonic/gin"
)

// Index category 列表
func Index(c *gin.Context) {
	cats, _ := categoryModel.All()

	controllers.SendOKResponse(c, controllers.ListData{
		List: viewmodels.CategoryList(cats),
	})
}
