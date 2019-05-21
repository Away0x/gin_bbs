package factory

import (
	"gin_bbs/app/models/permission"
)

// PermissionTableSeeder -
func PermissionTableSeeder(needCleanTable bool) {
	if needCleanTable {
		dropAndCreateTable(&permission.Permission{})
		dropAndCreateTable(&permission.Role{})
		dropAndCreateTable(&permission.ModelHasRole{})
		dropAndCreateTable(&permission.RoleHasPermission{})
		dropAndCreateTable(&permission.ModelHasPermission{})
	}

	// 创建权限
	ps := []*permission.Permission{
		{GuardName: permission.GuardNameWeb, Name: permission.PermissionNameManageContents},
		{GuardName: permission.GuardNameWeb, Name: permission.PermissionNameManageUsers},
		{GuardName: permission.GuardNameWeb, Name: permission.PermissionNameEditSettings},
	}
	for _, v := range ps {
		if err := v.Create(); err != nil {
			panic(err)
		}
	}

	// 创建站长角色，并赋予权限
	founder := &permission.Role{GuardName: permission.GuardNameWeb, Name: permission.RoleNameFounder}
	if err := founder.Create(); err != nil {
		panic(err)
	}
	for _, v := range ps {
		founder.GivePermissionTo(v.Name)
	}

	// 创建管理员角色，并赋予权限
	maintainer := &permission.Role{GuardName: permission.GuardNameWeb, Name: permission.RoleNameMaintainer}
	if err := maintainer.Create(); err != nil {
		panic(err)
	}
	maintainer.GivePermissionTo(permission.PermissionNameManageContents)
}
