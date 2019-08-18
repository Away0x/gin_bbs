package factory

import (
	"fmt"
	"gin_bbs/app/models"
	permissionModel "gin_bbs/app/models/permission"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/pkg/ginutils/utils"

	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

var (
	// 头像假数据
	avatars = []string{
		"https://cdn.learnku.com/uploads/avatars/7850_1481780622.jpeg!/both/380x380",
		"https://cdn.learnku.com/uploads/avatars/7850_1481780622.jpeg!/both/380x380",
		"https://cdn.learnku.com/uploads/avatars/7850_1481780622.jpeg!/both/380x380",
		"https://cdn.learnku.com/uploads/avatars/7850_1481780622.jpeg!/both/380x380",
		"https://cdn.learnku.com/uploads/avatars/7850_1481780622.jpeg!/both/380x380",
		"https://cdn.learnku.com/uploads/avatars/7850_1481780622.jpeg!/both/380x380",
	}
)

func userFactory(i int) *factory.Factory {
	now := time.Now()
	u := &userModel.User{
		Password:        "123456",
		EmailVerifiedAt: &now,
		Activated:       models.TrueTinyint,
		RememberToken:   string(utils.RandomCreateBytes(10)),
	}
	r := utils.RandInt(0, len(avatars)-1)

	return factory.NewFactory(
		u,
	).Attr("Name", func(args factory.Args) (interface{}, error) {
		return fmt.Sprintf("user-%d", i+1), nil
	}).Attr("Avatar", func(args factory.Args) (interface{}, error) {
		return avatars[r], nil
	}).Attr("Introduction", func(args factory.Args) (interface{}, error) {
		paragraph := randomdata.Paragraph()

		if len(paragraph) >= 70 {
			paragraph = paragraph[:70]
		}
		return paragraph, nil
	}).Attr("Email", func(args factory.Args) (interface{}, error) {
		if i == 0 {
			return "1@test.com", nil
		}
		if i == 1 {
			return "2@test.com", nil
		}
		return randomdata.Email(), nil
	})
}

// UsersTableSeeder -
func UsersTableSeeder(needCleanTable bool) {
	if needCleanTable {
		dropAndCreateTable(&userModel.User{})
	}

	for i := 0; i < 10; i++ {
		user := userFactory(i).MustCreate().(*userModel.User)
		if err := user.Create(); err != nil {
			fmt.Printf("mock user error： %v\n", err)
		}

		// 1号用户指派为 "站长"
		if i == 0 {
			r, err := permissionModel.GetRoleByName(permissionModel.RoleNameFounder)
			if err != nil {
				panic(err)
			}
			r.AssignRole(user)
		}
		// 2号用户指派为 "管理员"
		if i == 1 {
			r, err := permissionModel.GetRoleByName(permissionModel.RoleNameMaintainer)
			if err != nil {
				panic(err)
			}
			r.AssignRole(user)
		}
	}

}
