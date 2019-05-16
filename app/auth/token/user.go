package token

// // GetTokenFromRequest 从请求中获取 token
// func GetTokenFromRequest(c *gin.Context) (string, *errno.Errno) {
// 	header := c.Request.Header.Get("Authorization")
// 	if header == "" {
// 		return "", errno.TokenMissingError
// 	}

// 	var token string
// 	fmt.Sscanf(header, tokenIdentification+" %s", &token)
// 	return token, nil
// }

// // ParseAndGetUser 解析 token 获取 user
// func ParseAndGetUser(token string) (*userModel.User, *errno.Errno) {
// 	claims, err := parse(token)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// 判断是否过期
// 	exp := claims["exp"].(int64)
// 	if isExpire(exp) {
// 		return nil, errno.TokenExpireError
// 	}

// 	if id, ok := claims["userid"]; ok {
// 		intid := id.(int)
// 		u, err := userModel.Get(intid)
// 		if err != nil {
// 			return nil, errno.DatabaseError
// 		}
// 		return u, nil
// 	}

// 	return nil, errno.TokenError
// }
