package viewmodels

import (
	permissionModel "gin_bbs/app/models/permission"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/pkg/constants"
	gintime "gin_bbs/pkg/ginutils/time"
)

// UserViewModel 用户
type UserViewModel struct {
	ID                int
	Name              string
	Email             string
	Avatar            string
	Introduction      string
	CreatedAt         string
	NotificationCount int
}

// NewUserViewModelSerializer 用户数据展示
func NewUserViewModelSerializer(u *userModel.User) *UserViewModel {
	return &UserViewModel{
		ID:                int(u.ID),
		Name:              u.Name,
		Email:             u.Email,
		Avatar:            u.Avatar,
		Introduction:      u.Introduction,
		NotificationCount: u.NotificationCount,
		CreatedAt:         gintime.SinceForHuman(u.CreatedAt),
	}
}

// NewUserAPISerializer api data
func NewUserAPISerializer(u *userModel.User) map[string]interface{} {
	r := map[string]interface{}{
		"id":              u.ID,
		"name":            u.Name,
		"email":           u.Email,
		"avatar":          u.Avatar,
		"introduction":    u.Introduction,
		"bound_phone":     false,
		"bound_wechat":    false,
		"last_actived_at": u.EmailVerifiedAt.Format(constants.DateTimeLayout),
		"created_at":      u.CreatedAt.Format(constants.DateTimeLayout),
		"updated_at":      u.UpdatedAt.Format(constants.DateTimeLayout),
	}

	if u.Phone != "" {
		r["bound_phone"] = true
	}
	if u.WeixinUnionID != "" {
		r["bound_wechat"] = true
	}

	return r
}

// NewUserAPIHasRoles -
func NewUserAPIHasRoles(u *userModel.User, rs []*permissionModel.Role) map[string]interface{} {
	uvm := NewUserAPISerializer(u)
	uvm["roles"] = RoleAPIList(rs)

	return uvm
}
