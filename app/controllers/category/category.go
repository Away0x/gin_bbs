package category

import (
	"gin_bbs/app/controllers"
	categoryModel "gin_bbs/app/models/category"
	topicModel "gin_bbs/app/models/topic"
	"gin_bbs/app/services"
	"gin_bbs/pkg/ginutils"
	"gin_bbs/pkg/ginutils/pagination"

	"github.com/gin-gonic/gin"
)

// Show topic 详情
func Show(c *gin.Context) {
	id, err := ginutils.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	cat, err := categoryModel.Get(id)
	if err != nil {
		controllers.Render404(c)
		return
	}

	renderFunc, err := pagination.CreatePage(c, 20, "topics",
		func() (int, error) { return topicModel.CountByCategoryID(int(cat.ID)) },
		func(offset, limit, _, _ int) (interface{}, error) {
			return services.TopicListService(func() ([]*topicModel.Topic, error) {
				return topicModel.GetByCategoryID(int(cat.ID), offset, limit, c.DefaultQuery("order", "default"))
			})
		})

	if err != nil {
		controllers.Render(c, "topics/index", gin.H{"error": "错误: " + err.Error()})
		return
	}

	controllers.Render(c, "topics/index", renderFunc(gin.H{
		"category": cat,
	}))
}
