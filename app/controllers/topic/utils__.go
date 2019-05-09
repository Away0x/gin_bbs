package topic

import (
	"gin_bbs/app/controllers"
	categoryModel "gin_bbs/app/models/category"
	topicModel "gin_bbs/app/models/topic"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/app/policies"
	"gin_bbs/pkg/ginutils"
	"gin_bbs/pkg/ginutils/utils"
	"gin_bbs/pkg/ginutils/validate"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取要编辑的 topic
func getEditTopic(c *gin.Context, currentUser *userModel.User) (*topicModel.Topic, int, bool) {
	id, err := ginutils.GetIntParam(c, "id")
	if err != nil {
		controllers.Render404(c)
		return nil, id, false
	}

	topic, err := topicModel.Get(id)
	if err != nil {
		controllers.Render404(c)
		return nil, id, false
	}

	// 权限
	if ok := policies.TopicPolicyOwner(c, currentUser, int(topic.UserID)); !ok {
		return nil, id, false
	}

	return topic, id, true
}

// 创建、更新 topic 时的参数验证
func validateCreateOrUpdateTopic(c *gin.Context, currentUser *userModel.User, topic *topicModel.Topic) bool {
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
			"title":       {"文章标题不能为空", "文章标题长度必须大于 2 个字符"},
			"category_id": {"文章分类不能为空", "选择的文章分类不存在"},
			"body":        {"文章内容不能为空", "文章内容长度必须大于 3 个字符"},
		})

	if !ok {
		validate.SaveValidateMessage(c, errArr, errMap)
		return false
	}

	topic.Title = title
	topic.Body = utils.XSSClean(body)
	topic.CategoryID = uint(catID)
	topic.UserID = currentUser.ID

	return true
}
