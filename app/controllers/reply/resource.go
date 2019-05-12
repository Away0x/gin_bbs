package reply

import (
	replyModel "gin_bbs/app/models/reply"
	topicModel "gin_bbs/app/models/topic"
	userModel "gin_bbs/app/models/user"

	"gin_bbs/app/controllers"
	"gin_bbs/pkg/ginutils/validate"
	"strconv"

	"gin_bbs/pkg/ginutils/flash"

	"github.com/gin-gonic/gin"
)

// Store 发表回复
func Store(c *gin.Context, currentUser *userModel.User) {
	var topic *topicModel.Topic
	content := c.PostForm("content")
	topicID := c.PostForm("topic_id")

	ok, errArr, errMap := validate.RunByParams(false,
		validate.ValidatorMap{
			"content": {
				validate.RequiredValidator(content),
				validate.MixLengthValidator(content, 2),
			},
			"topic_id": {
				func() string {
					id, err := strconv.Atoi(topicID)
					if err != nil {
						return "评论失败: " + err.Error()
					}
					topic, err = topicModel.Get(id)
					if err != nil {
						return "评论失败: " + err.Error()
					}
					return ""
				},
			},
		},
		validate.MessagesMap{"content": {"评论内容不能为空", "评论内容长度不得小于 2 个字符"}})

	if !ok {
		validate.SaveValidateMessage(c, errArr, errMap)
		if topic != nil {
			controllers.Redirect(c, topic.Link(), false)
			return
		}
		controllers.Render404(c)
		return
	}

	reply := &replyModel.Reply{
		TopicID: topic.ID,
		UserID:  currentUser.ID,
		Content: content,
	}
	if err := reply.Create(); err != nil {
		flash.NewDangerFlash(c, "评论创建失败: "+err.Error())
		controllers.Redirect(c, topic.Link(), false)
		return
	}

	flash.NewSuccessFlash(c, "评论创建成功")
	controllers.Redirect(c, topic.Link(), false)
}

// Destroy 删除回复
func Destroy(c *gin.Context, currentUser *userModel.User) {

}
