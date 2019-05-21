package viewmodels

import (
	permissionModel "gin_bbs/app/models/permission"
)

// PermissionAPI -
func PermissionAPI(p *permissionModel.Permission) map[string]interface{} {
	return map[string]interface{}{
		"id":   p.ID,
		"name": p.Name,
	}
}

// PermissionAPIList -
func PermissionAPIList(ps []*permissionModel.Permission) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	for _, v := range ps {
		result = append(result, PermissionAPI(v))
	}

	return result
}

// RoleAPI -
func RoleAPI(p *permissionModel.Role) map[string]interface{} {
	return map[string]interface{}{
		"id":   p.ID,
		"name": p.Name,
	}
}

// RoleAPIList -
func RoleAPIList(ps []*permissionModel.Role) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	for _, v := range ps {
		result = append(result, RoleAPI(v))
	}

	return result
}
