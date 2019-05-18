package viewmodels

import (
	categoryModel "gin_bbs/app/models/category"
)

// Category -
func Category(c *categoryModel.Category) map[string]interface{} {
	return map[string]interface{}{
		"id":          c.ID,
		"name":        c.Name,
		"description": c.Description,
	}
}

// CategoryList -
func CategoryList(cs []*categoryModel.Category) []map[string]interface{} {
	vs := make([]map[string]interface{}, 0)
	for _, v := range cs {
		vs = append(vs, Category(v))
	}

	return vs
}
