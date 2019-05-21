package viewmodels

import (
	linkModel "gin_bbs/app/models/link"
)

// LinkAPI -
func LinkAPI(l *linkModel.Link) map[string]interface{} {
	return map[string]interface{}{
		"id":    l.ID,
		"title": l.Title,
		"link":  l.Link,
	}
}

// LinkAPIList -
func LinkAPIList(ls []*linkModel.Link) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	for _, v := range ls {
		result = append(result, LinkAPI(v))
	}

	return result
}
