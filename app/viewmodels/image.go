package viewmodels

import (
	imageModel "gin_bbs/app/models/image"
	"gin_bbs/pkg/constants"
)

// Image -
func Image(i *imageModel.Image) map[string]interface{} {
	return map[string]interface{}{
		"id":         i.ID,
		"user_id":    i.UserID,
		"type":       i.Type,
		"path":       i.Path,
		"created_at": i.CreatedAt.Format(constants.DateTimeLayout),
		"updated_at": i.UpdatedAt.Format(constants.DateTimeLayout),
	}
}
