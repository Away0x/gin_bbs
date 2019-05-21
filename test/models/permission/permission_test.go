package permission

import (
	permissionModel "gin_bbs/app/models/permission"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/bootstrap"
	"gin_bbs/config"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPermissionModel(t *testing.T) {
	config.InitConfig("./../../../config.yaml", false)
	// db setup
	db, err := bootstrap.SetupDB(false)
	if err != nil {
		return
	}
	defer db.Close()

	var ok bool
	founder, _ := userModel.Get(1)
	maintainer, _ := userModel.Get(2)

	Convey("UserHasRole", t, func() {
		isFounder := permissionModel.UserHasRole(founder, permissionModel.RoleNameFounder)
		So(isFounder, ShouldEqual, true)
	})

	Convey("UserHasRole", t, func() {
		isMaintainer := permissionModel.UserHasRole(maintainer, permissionModel.RoleNameMaintainer)
		So(isMaintainer, ShouldEqual, true)
	})

	Convey("UserHasPermission", t, func() {
		ok = permissionModel.UserHasPermission(founder, permissionModel.PermissionNameManageContents)
		So(ok, ShouldEqual, true)

		ok = permissionModel.UserHasPermission(maintainer, permissionModel.PermissionNameManageUsers)
		So(ok, ShouldEqual, false)

		ok = permissionModel.UserHasPermission(founder, "xxxx")
		So(ok, ShouldEqual, false)
	})

	Convey("GetUserAllPermission", t, func() {
		ps, _ := permissionModel.GetUserAllPermission(founder)
		So(len(ps), ShouldEqual, 3)

		ps, _ = permissionModel.GetUserAllPermission(maintainer)
		So(len(ps), ShouldEqual, 1)
	})
}
