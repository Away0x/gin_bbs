package topic

import (
	"gin_bbs/app/controllers"
	"gin_bbs/pkg/ginutils/pagination"

	topicModel "gin_bbs/app/models/topic"

	"gin_bbs/app/services"

	"github.com/gin-gonic/gin"
)

// Index topic 列表
func Index(c *gin.Context) {
	renderFunc, err := pagination.CreatePage(c, 30, "topics", topicModel.Count,
		func(offset, limit, _, _ int) (interface{}, error) {
			return services.TopicListService(func() ([]*topicModel.Topic, error) {
				return topicModel.List(offset, limit)
			})
		})

	if err != nil {
		controllers.Render(c, "topics/index", gin.H{"error": "错误: " + err.Error()})
		return
	}

	controllers.Render(c, "topics/index", renderFunc(gin.H{}))
}

// Show topic 详情
func Show(c *gin.Context) {

}

// // Create 创建 topic 页
// func Create(c *gin.Context) {

// }

// // Store 保存新 topic
// func Store(c *gin.Context) {

// }

// // Edit 编辑 topic 页面
// func Edit(c *gin.Context) {

// }

// // Update 编辑 topic
// func Update(c *gin.Context) {

// }

// // Destroy 删除 topic
// func Destroy(c *gin.Context) {

// }
