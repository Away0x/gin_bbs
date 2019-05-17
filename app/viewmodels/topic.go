package viewmodels

import (
	topicModel "gin_bbs/app/models/topic"
	"gin_bbs/pkg/constants"
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
		"Link":            t.Link(),
	}
}

// Topic -
func TopicApi(t *topicModel.Topic) map[string]interface{} {
	return map[string]interface{}{
		"id":                 t.ID,
		"title":              t.Title,
		"body":               t.Body,
		"user_id":            t.UserID,
		"category_id":        t.CategoryID,
		"reply_count":        t.ReplyCount,
		"view_count":         t.ViewCount,
		"last_reply_user_id": t.LastReplyUserID,
		"excerpt":            t.Excerpt,
		"slug":               t.Slug,
		"created_at":         t.CreatedAt.Format(constants.DateTimeLayout),
		"updated_at":         t.UpdatedAt.Format(constants.DateTimeLayout),
	}
}
