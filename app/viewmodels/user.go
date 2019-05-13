package viewmodels

import (
	userModel "gin_bbs/app/models/user"
	gintime "gin_bbs/pkg/ginutils/time"
)

// UserViewModel 用户
type UserViewModel struct {
	ID           int
	Name         string
	Email        string
	Avatar       string
	Introduction string
	CreatedAt     string
	NotificationCount int
}

// NewUserViewModelSerializer 用户数据展示
func NewUserViewModelSerializer(u *userModel.User) *UserViewModel {
	return &UserViewModel{
		ID:           int(u.ID),
		Name:         u.Name,
		Email:        u.Email,
		Avatar:       u.Avatar,
		Introduction: u.Introduction,
		NotificationCount: u.NotificationCount,
		CreatedAt:     gintime.SinceForHuman(u.CreatedAt),
	}
}
