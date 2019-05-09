package viewmodels

import (
	topicModel "gin_bbs/app/models/topic"
	gintime "gin_bbs/pkg/ginutils/time"
)

// NewTopicViewModelSerializer -
func NewTopicViewModelSerializer(t *topicModel.Topic) map[string]interface{} {
	return map[string]interface{}{
		"ID":              t.ID,
		"CategoryID":      t.CategoryID,
		"UserID":          t.UserID,
		"Title":           t.Title,
		"Body":            t.Body,
		"ReplyCount":      t.ReplyCount,
		"ViewCount":       t.ViewCount,
		"LastReplyUserID": t.LastReplyUserID,
		"Order":           t.Order,
		"Excerpt":         t.Excerpt,
		"Slug":            t.Slug,
		"CreatedAt":       gintime.SinceForHuman(t.CreatedAt),
		"UpdatedAt":       gintime.SinceForHuman(t.UpdatedAt),
	}
}
