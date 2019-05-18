package reply

import (
	replyModel "gin_bbs/app/models/reply"
	topicModel "gin_bbs/app/models/topic"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/app/viewmodels"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils"
	"gin_bbs/pkg/ginutils/validate"

	"gin_bbs/app/controllers"
	"gin_bbs/app/policies"

	"github.com/gin-gonic/gin"
)

type storeParams struct {
	Content string `json:"content"`
	TopicID uint   `json:"topic_id"`
}

// Store 发表回复
func Store(c *gin.Context, currentUser *userModel.User, tokenString string) {
	var (
		req   storeParams
		topic *topicModel.Topic
		err   error
	)
	if err = c.ShouldBind(&req); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}

	ok, _, msgMap := validate.RunByParams(false,
		validate.ValidatorMap{
			"content": {
				validate.RequiredValidator(req.Content),
				validate.MinLengthValidator(req.Content, 2),
			},
			"topic_id": {
				func() string {
					if req.TopicID == 0 {
						return "话题 ID 不可为空"
					}

					topic, err = topicModel.Get(int(req.TopicID))
					if err != nil {
						return "话题不存在"
					}

					return ""
				},
			},
		},
		validate.MessagesMap{"content": {"回复内容不能为空", "回复内容不能小于 2 个字符"}})
	if !ok {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, msgMap))
		return
	}

	// 创建 reply
	reply := &replyModel.Reply{
		TopicID: topic.ID,
		UserID:  currentUser.ID,
		Content: req.Content,
	}
	if err = reply.Create(); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.DatabaseError, err))
		return
	}

	replyModel.TopicRepliedNotify(reply, currentUser) // 通知
	controllers.SendOKResponse(c, viewmodels.ReplyApi(reply))
}

// Destroy 删除回复
func Destroy(c *gin.Context, currentUser *userModel.User, tokenString string) {
	replyID, err := ginutils.GetIntParam(c, "id")
	if err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}

	reply, err := replyModel.Get(replyID)
	if err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ResourceNotFoundError, err))
		return
	}

	topic, err := topicModel.Get(int(reply.TopicID))
	if err != nil {
		controllers.Render404(c)
		return
	}

	// 权限
	if !policies.CheckPolicy(c, func() bool {
		return reply.UserID == currentUser.ID || topic.UserID == currentUser.ID
	}) {
		return
	}

	if err := replyModel.Delete(replyID); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.DatabaseError, err))
		return
	}

	controllers.SendOKResponse(c, map[string]interface{}{
		"id": replyID,
	})
}
