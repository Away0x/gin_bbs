package token

import (
	userModel "gin_bbs/app/models/user"
	"gin_bbs/pkg/constants"
	"gin_bbs/pkg/errno"
	"time"
)

// Info token info
type Info struct {
	Token     string `json:"token"`
	Type      string `json:"type"`
	ExpiresIn string `json:"expires_in"`
}

// Sign 签发 token
func Sign(u *userModel.User) (*Info, *errno.Errno) {
	t, claims, err := create(u)
	if err != nil || t == "" {
		return nil, errno.New(errno.TokenError, err)
	}

	return &Info{
		Token:     t,
		Type:      tokenInHeaderIdentification,
		ExpiresIn: time.Unix(claims.ExpiresAt, 0).Format(constants.DateTimeLayout),
	}, nil
}

// Refresh 刷新 token
func Refresh(tokenString string) (*Info, *errno.Errno) {
	t, claims, err := refresh(tokenString)
	if err != nil || t == "" {
		return nil, err
	}

	return &Info{
		Token:     t,
		Type:      tokenInHeaderIdentification,
		ExpiresIn: time.Unix(claims.ExpiresAt, 0).Format(constants.DateTimeLayout),
	}, nil
}

// Forget 使 token 失效
func Forget(tokenString string, remainTime time.Duration) {
	forget(tokenString, remainTime)
}
