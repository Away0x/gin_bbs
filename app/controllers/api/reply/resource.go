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

	"github.com/gin-gonic/gin"
)

type storeParams struct {
	Content string `json:"content"`
}

// Store 发表回复
func Store(c *gin.Context, currentUser *userModel.User, tokenString string) {
	// 参数校验
	topicID, err := ginutils.GetIntParam(c, "topic_id")
	if err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}

	var req storeParams
	if err := c.ShouldBind(&req); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}
	ok, _, msgMap := validate.RunSingle("content",
		[]validate.ValidatorFunc{
			validate.RequiredValidator(req.Content),
			validate.MinLengthValidator(req.Content, 2),
		},
		[]string{"回复内容不能为空", "回复内容不能小于 2 个字符"})
	if !ok {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, msgMap))
		return
	}

	// 获取 topic
	topic, err := topicModel.Get(topicID)
	if err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.DatabaseError, err))
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

	controllers.SendOKResponse(c, viewmodels.ReplyApi(reply))
}
