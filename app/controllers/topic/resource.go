package topic

import (
	"gin_bbs/app/controllers"
	"gin_bbs/pkg/ginutils"
	"gin_bbs/pkg/ginutils/pagination"
	"strings"

	categoryModel "gin_bbs/app/models/category"
	replyModel "gin_bbs/app/models/reply"
	topicModel "gin_bbs/app/models/topic"
	userModel "gin_bbs/app/models/user"

	"gin_bbs/app/auth"
	"gin_bbs/app/policies"
	"gin_bbs/app/services"
	"gin_bbs/app/viewmodels"
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

	// url 矫正 (永远带 slug 参数来访问)
	ReqSlug := c.Param("slug")
	ReqSlug = strings.Replace(ReqSlug, "/", "", -1) // ReqSlug 参数可能这样 /slug
	if ReqSlug == "" && topic.Slug != ReqSlug {
		c.Redirect(301, topic.Link())
		return
	}

	topicVM := viewmodels.NewTopicViewModelSerializer(topic)
	topicVM["User"] = viewmodels.NewUserViewModelSerializer(user)

	// 获取回复
	currentUser, _ := auth.GetCurrentUserFromContext(c) // 当前用户
	replies, _ := services.RpleyListService(func() ([]*replyModel.Reply, error) {
		return replyModel.TopicReplies(int(topic.ID))
	}, currentUser)

	controllers.Render(c, "topics/show", gin.H{
		"topic":          topicVM,
		"replies":        replies,
		"canDeleteTopic": policies.CheckTopicPolicyOwner(currentUser, int(topic.UserID)),
	})
}

// Create 创建 topic 页
func Create(c *gin.Context, currentUser *userModel.User) {
	categories, _ := categoryModel.All()
	controllers.Render(c, "topics/create_and_edit", gin.H{"categories": categories})
}

// Store 保存新 topic
func Store(c *gin.Context, currentUser *userModel.User) {
	topic := &topicModel.Topic{}
	if ok := validateCreateOrUpdateTopic(c, currentUser, topic); !ok {
		controllers.RedirectRouter(c, "topics.create")
		return
	}

	if err := topic.Create(); err != nil {
		flash.NewDangerFlash(c, "帖子创建失败: "+err.Error())
		controllers.RedirectRouter(c, "topics.create")
		return
	}

	flash.NewSuccessFlash(c, "帖子创建成功")
	controllers.Redirect(c, topic.Link(), false)
}

// Edit 编辑 topic 页面
func Edit(c *gin.Context, currentUser *userModel.User) {
	topic, _, ok := getEditTopic(c, currentUser)
	if !ok {
		return
	}

	categories, _ := categoryModel.All()
	controllers.Render(c, "topics/create_and_edit", gin.H{
		"categories": categories,
		"topic":      viewmodels.NewTopicViewModelSerializer(topic),
	})
}

// Update 编辑 topic
func Update(c *gin.Context, currentUser *userModel.User) {
	// 通过 param 中 id 参数获取 topic (没获取到或没权限会进入 404 page)
	topic, id, ok := getEditTopic(c, currentUser)
	if !ok {
		return
	}
	// 验证参数
	ok = validateCreateOrUpdateTopic(c, currentUser, topic)
	if !ok {
		controllers.RedirectRouter(c, "topics.edit", id)
		return
	}

	topic.ID = uint(id)
	if err := topic.Update(); err != nil {
		flash.NewDangerFlash(c, "帖子编辑失败: "+err.Error())
		controllers.RedirectRouter(c, "topics.edit", topic.ID)
		return
	}

	flash.NewSuccessFlash(c, "帖子编辑成功")
	controllers.Redirect(c, topic.Link(), false)
}

// Destroy 删除 topic
func Destroy(c *gin.Context, currentUser *userModel.User) {
	topic, id, ok := getEditTopic(c, currentUser)
	if !ok {
		return
	}

	if err := topicModel.Delete(id); err != nil {
		flash.NewDangerFlash(c, "帖子删除失败: "+err.Error())
		controllers.Redirect(c, topic.Link(), false)
		return
	}

	// 删除帖子的评论
	if err := replyModel.DeleteTopicReplies(id); err != nil {
		flash.NewDangerFlash(c, "评论删除失败: "+err.Error())
	}

	flash.NewSuccessFlash(c, "帖子删除成功")
	controllers.RedirectRouter(c, "topics.index")
}
