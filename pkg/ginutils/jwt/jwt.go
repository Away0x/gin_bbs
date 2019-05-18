package jwt

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type a struct {
	jwt.StandardClaims
}

// Sign 签发 token
func Sign(secret string, data map[string]interface{}) (tokenString string, err error) {
	if secret == "" {
		return "", errors.New("secret 不能为空")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	for k, v := range data {
		claims[k] = v
	}

	token.Claims = claims
	tokenString, err = token.SignedString([]byte(secret))

	return
}

// Parse 解析 token
func Parse(secret string, tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("jwt token parse error")
}
