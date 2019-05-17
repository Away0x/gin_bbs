package token

import (
	"fmt"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/pkg/errno"

	"github.com/gin-gonic/gin"
)

const (
	tokenHeaderKeyName          = "Authorization"
	tokenInHeaderIdentification = "Bearer"
)

// GetTokenFromRequest 从请求中获取 token
func GetTokenFromRequest(c *gin.Context) (string, *errno.Errno) {
	header := c.Request.Header.Get(tokenHeaderKeyName)
	if header == "" {
		return "", errno.TokenMissingError
	}

	var token string
	fmt.Sscanf(header, tokenInHeaderIdentification+" %s", &token)
	return token, nil
}

// ParseAndGetUser 解析 token 获取 user
func ParseAndGetUser(c *gin.Context, token string) (*userModel.User, *errno.Errno) {
	claims, err := parseToken(token)
	if err != nil {
		return nil, err
	}

	user, e := userModel.Get(int(claims.UserID))
	if e != nil {
		return nil, errno.New(errno.DatabaseError, e)
	}

	c.Set(tokenHeaderKeyName+"User", user)
	c.Set(tokenHeaderKeyName+"Token", token)
	return user, nil
}

// GetTokenUserFromContext -
func GetTokenUserFromContext(c *gin.Context) (string, *userModel.User, bool) {
	user, ok := c.Get(tokenHeaderKeyName + "User")
	if !ok {
		return "", nil, false
	}
	t, ok := c.Get(tokenHeaderKeyName + "Token")
	if !ok {
		return "", nil, false
	}

	u, ok := user.(*userModel.User)
	s, ok := t.(string)
	if !ok {
		return "", nil, false
	}

	return s, u, true
}
