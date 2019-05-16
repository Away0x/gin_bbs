package token

import (
	"errors"
	"gin_bbs/app/cache"
	"gin_bbs/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	cacheTokenKeyName   = "jwt_token_"
	jwtTokenRefreshTime = 20160 * time.Minute // 允许刷新 token 的时间 (14 天)
	jwtTokenExpiredTime = 1 * time.Minute     // token 60 分钟过期
	tokenIdentification = "Bearer"
)

// CustomClaims -
type CustomClaims struct {
	jwt.StandardClaims
	UserID      uint  `json:"userid"`
	RefreshTime int64 `json:"refresh_time,omitempty"`
}

// SetExpiredAt 设置 token 有效期
func (c *CustomClaims) SetExpiredAt() {
	now := time.Now()
	c.IssuedAt = now.Unix()
	c.ExpiresAt = now.Add(jwtTokenExpiredTime).Unix()
	c.RefreshTime = now.Add(jwtTokenRefreshTime).Unix()
}

// Create 创建 token
func Create(userid uint) (string, CustomClaims, error) {
	if config.AppConfig.Key == "" {
		return "", CustomClaims{}, errors.New("jwt secret 不能为空")
	}

	claims := CustomClaims{
		UserID: userid,
	}
	claims.SetExpiredAt()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(config.AppConfig.Key))
	if err != nil {
		return "", claims, err
	}
	return s, claims, err
}

// Parse 解析 token
func Parse(tokenString string) (*CustomClaims, error) {
	token, err := parse(tokenString)
	if err != nil {
		if isExpired(err) {
			if claims, ok := token.Claims.(*CustomClaims); ok {
				return claims, errors.New("token 过期了")
			}
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("jwt token parse error")
}

// Refresh 刷新 token
func Refresh(tokenString string, userid uint) (string, CustomClaims, error) {
	token, err := parse(tokenString)
	if err != nil {
		// 非过期错误
		if !isExpired(err) {
			return "", CustomClaims{}, err
		}
		// 判断是否过了可刷新时间
		if claims, ok := token.Claims.(*CustomClaims); ok {
			now := time.Now().Unix()
			if now > claims.RefreshTime {
				return "", CustomClaims{}, errors.New("token 过了刷新时间了")
			}
		}
	}

	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return "", CustomClaims{}, err
	}

	Forget(tokenString) // 将之前的 token 加入黑名单使之失效
	newToken, newClaims, err := Create(claims.UserID)
	if err != nil {
		return "", CustomClaims{}, err
	}

	return newToken, newClaims, nil
}

// Forget 使 token 失效
func Forget(tokenString string) {
	key := cacheTokenKeyName + tokenString
	cache.Put(key, "1", jwtTokenExpiredTime)
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

func parse(tokenString string) (token *jwt.Token, err error) {
	token, err = jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.Key), nil
	})

	return
}
