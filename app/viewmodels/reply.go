package viewmodels

import (
	replyModel "gin_bbs/app/models/reply"
	gintime "gin_bbs/pkg/ginutils/time"
)

// NewReplyViewModelSerializer -
func NewReplyViewModelSerializer(r *replyModel.Reply) map[string]interface{} {
	return map[string]interface{}{
			"ID": r.ID,
			"Content": r.Content,
			"UserID": r.UserID,
			"TopicID": r.TopicID,
			"CreatedAt": gintime.SinceForHuman(r.CreatedAt),
	}
}
