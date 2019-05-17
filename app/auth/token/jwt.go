package token

import (
	"gin_bbs/app/cache"
	userModel "gin_bbs/app/models/user"
	"gin_bbs/config"
	"gin_bbs/pkg/errno"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	cacheTokenKeyName   = "jwt_token_"
	jwtTokenRefreshTime = 20160 * time.Minute // 允许刷新 token 的时间 (14 天); 期间内允许使用之前颁发的 token (即使过期)来获取新token
	jwtTokenExpiredTime = 60 * time.Minute    // token 60 分钟过期
	jwtTokenRemainTime  = 2 * time.Minute     // token 刷新后，旧的 token 还有 2 分钟的使用时间 (前提是旧 token 没过过期时间)
)

func getCacheKey(tokenString string) string {
	return cacheTokenKeyName + tokenString
}

// CustomClaims -
type CustomClaims struct {
	jwt.StandardClaims
	UserID      uint  `json:"userid"`
	RefreshTime int64 `json:"refresh_time,omitempty"`
}

// SetUser 设置 token 有效期
func (c *CustomClaims) SetUser(u *userModel.User) {
	c.UserID = u.ID
}

// SetExpiredAt 设置 user data
func (c *CustomClaims) SetExpiredAt() {
	now := time.Now()
	c.IssuedAt = now.Unix()
	c.ExpiresAt = now.Add(jwtTokenExpiredTime).Unix()
	c.RefreshTime = now.Add(jwtTokenRefreshTime).Unix()
}

// create 创建 token
func create(u *userModel.User) (string, CustomClaims, error) {
	claims := CustomClaims{}
	claims.SetUser(u)
	claims.SetExpiredAt()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(config.AppConfig.Key))
	if err != nil {
		return "", claims, err
	}
	return s, claims, nil
}

// parseToken 解析 token
func parseToken(tokenString string) (*CustomClaims, *errno.Errno) {
	token, err := parse(tokenString)
	if err != nil {
		// token 过期
		if isExpired(err) {
			if claims, ok := token.Claims.(*CustomClaims); ok {
				return claims, errno.TokenExpireError
			}
		}
		if e, ok := err.(*errno.Errno); ok {
			return nil, e
		}
		return nil, errno.New(errno.TokenError, err)
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errno.New(errno.TokenError, "jwt token parse error")
}

// refresh 刷新 token
func refresh(tokenString string) (string, CustomClaims, *errno.Errno) {
	token, err := parse(tokenString)
	if err != nil {
		// 非过期错误
		if !isExpired(err) {
			return "", CustomClaims{}, errno.New(errno.TokenError, err)
		}
		// 判断是否过了可刷新时间
		if claims, ok := token.Claims.(*CustomClaims); ok {
			now := time.Now().Unix()
			if now > claims.RefreshTime {
				return "", CustomClaims{}, errno.TokenExpireError
			}
		}
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return "", CustomClaims{}, errno.TokenError
	}

	forget(tokenString, jwtTokenRemainTime) // 将之前的 token 加入黑名单使之失效
	u := &userModel.User{}
	u.ID = claims.UserID
	newToken, newClaims, err := create(u)
	if err != nil {
		return "", CustomClaims{}, errno.New(errno.TokenError, err)
	}

	return newToken, newClaims, nil
}

// forget 使 token 失效
func forget(tokenString string, remainTime time.Duration) {
	now := time.Now()
	cache.PutInt64(getCacheKey(tokenString), now.Add(remainTime).Unix(), jwtTokenExpiredTime) // val 为 token 还可以使用的过渡时间
}

// ------------- private
func isExpired(err error) bool {
	switch err.(type) {
	case *jwt.ValidationError:
		vErr := err.(*jwt.ValidationError)
		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			// token 过期了
			return true
		default:
			return false
		}

	default:
		return false
	}
}

func parse(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.Key), nil
	})

	// 在黑名单中
	if t, ok := cache.GetInt64(getCacheKey(tokenString)); ok {
		now := time.Now().Unix()
		// 过了留存时间了
		if now > t {
			return nil, errno.TokenInBlackListError
		}
	}

	return token, err
}
