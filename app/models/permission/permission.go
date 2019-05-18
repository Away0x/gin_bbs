package permission

// Role 角色的模型表
type Role struct {
}

// TableName 表名
func (Role) TableName() string {
	return "roles"
}

// Permission 权限的模型表
type Permission struct {
}

// TableName 表名
func (Permission) TableName() string {
	return "permissions"
}

// ModelHasRole 模型与角色的关联表，用户拥有什么角色在此表中定义，一个用户能拥有多个角色
type ModelHasRole struct {
}

// TableName 表名
func (ModelHasRole) TableName() string {
	return "model_has_roles"
}

// RoleHasPermission 角色拥有的权限关联表，如管理员拥有查看后台的权限都是在此表定义，一个角色能拥有多个权限
type RoleHasPermission struct {
}

// TableName 表名
func (RoleHasPermission) TableName() string {
	return "role_has_permissions"
}

// ModelHasPermission 模型与权限关联表，一个模型能拥有多个权限
type ModelHasPermission struct {
}

// TableName 表名
func (ModelHasPermission) TableName() string {
	return "model_has_permissions"
}
