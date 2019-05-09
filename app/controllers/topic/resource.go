package topic

import (
	"gin_bbs/app/controllers"
	"gin_bbs/pkg/ginutils/pagination"
	"gin_bbs/pkg/ginutils"
	"strconv"

	categoryModel "gin_bbs/app/models/category"
	topicModel "gin_bbs/app/models/topic"
	userModel "gin_bbs/app/models/user"

	"gin_bbs/app/services"
	"gin_bbs/app/viewmodels"
	"gin_bbs/pkg/ginutils/validate"
	"gin_bbs/pkg/ginutils/flash"

	"github.com/gin-gonic/gin"
)

// Index topic 列表
func Index(c *gin.Context) {
	renderFunc, err := pagination.CreatePage(c, 20, "topics", topicModel.Count,
		func(offset, limit, _, _ int) (interface{}, error) {
			return services.TopicListService(func() ([]*topicModel.Topic, error) {
				return topicModel.List(offset, limit, c.DefaultQuery("order", "default"))
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
	id, err := ginutils.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return
	}

	topic, user, err := topicModel.TopicAndUser(id)
	if err != nil {
		controllers.Render404(c)
		return
	}

	topicVM := viewmodels.NewTopicViewModelSerializer(topic)
	topicVM["User"] = viewmodels.NewUserViewModelSerializer(user)

	controllers.Render(c, "topics/show", gin.H{
		"topic": topicVM,
	})
}

// Create 创建 topic 页
func Create(c *gin.Context, currentUser *userModel.User) {
	categories, _ := categoryModel.All()

	controllers.Render(c, "topics/create_and_edit", gin.H{
		"categories": categories,
	})
}

// Store 保存新 topic
func Store(c *gin.Context, currentUser *userModel.User) {
	title := c.PostForm("title")
	catIDStr := c.PostForm("category_id")
	catID, _ := strconv.Atoi(catIDStr)
	body := c.PostForm("body")
	categoryAllIDs, _ := categoryModel.AllID()

	ok, errArr, errMap := validate.RunByParams(false,
		validate.ValidatorMap{
			"title": {
				validate.RequiredValidator(title),
				validate.MixLengthValidator(title, 2),
			},
			"category_id": {
				validate.RequiredValidator(catIDStr),
				validate.UintRangeValidator(uint(catID), categoryAllIDs),
			},
			"body": {
				validate.RequiredValidator(body),
				validate.MixLengthValidator(body, 3),
			},
		},
		validate.MessagesMap{
			"title": {"文章标题不能为空", "文章标题长度必须大于 2 个字符"},
			"category_id": {"文章分类不能为空", "选择的文章分类不存在"},
			"body": {"文章内容不能为空", "文章内容长度必须大于 3 个字符"},
		})

	if !ok {
		validate.SaveValidateMessage(c, errArr, errMap)
		controllers.RedirectRouter(c, "topics.create")
		return
	}

	topic := &topicModel.Topic{
		Title: title,
		Body: body,
		CategoryID: uint(catID),
		UserID: currentUser.ID,
	}
	if err := topic.Create(); err != nil {
		flash.NewDangerFlash(c, "帖子创建失败: " + err.Error())
		controllers.RedirectRouter(c, "topics.create")
		return
	}

	flash.NewSuccessFlash(c, "帖子创建成功")
	controllers.RedirectRouter(c, "topics.show", topic.ID)
}

// // Edit 编辑 topic 页面
// func Edit(c *gin.Context) {

// }

// // Update 编辑 topic
// func Update(c *gin.Context) {

// }

// // Destroy 删除 topic
// func Destroy(c *gin.Context) {

// }
