package authorization

import (
	"gin_bbs/app/auth/token"
	"gin_bbs/app/controllers"
	userModel "gin_bbs/app/models/user"
	authorizationRequest "gin_bbs/app/requests/api/authorization"
	"gin_bbs/app/viewmodels"
	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils/utils"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	types = []string{"weixin"}
)

// SocialStore 第三方登录
// @Summary 第三方登录
// @Tags authorization
// @Accept  json
// @Produce  json
// @Param social_type path string true "social_type in [weixin]"
// @Param json body authorization.Social true "微信 access_token openid 和 code，要么传 access_token openid 要么只传 code"
// @Success 200 {object} controllers.Response "{"token": 1}"
// @Router /api/socials/authorizations/{social_type} [post]
func SocialStore(c *gin.Context) {
	var req *authorizationRequest.Social
	if err := c.ShouldBind(&req); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}
	req.SocialType = c.Param("social_type")
	weixinUserInfo, eno := req.Run()
	if eno != nil {
		controllers.SendErrorResponse(c, eno)
		return
	}

	var (
		user *userModel.User
		err  error
	)
	switch req.SocialType {
	case "weixin":
		if weixinUserInfo.Unionid != "" {
			user, err = userModel.GetByWeixinUnionID(weixinUserInfo.Unionid)
		} else {
			user, err = userModel.GetByWeixinOpenID(weixinUserInfo.OpenID)
		}

		if err != nil || user == nil {
			// 没有用户，默认创建一个用户
			user = &userModel.User{
				Name:         weixinUserInfo.NickName,
				Avatar:       weixinUserInfo.HeadImgURL,
				WeixinOpenID: weixinUserInfo.OpenID,
			}
			if weixinUserInfo.Unionid != "" {
				user.WeixinUnionID = weixinUserInfo.Unionid
			}

			if err := user.Create(); err != nil {
				controllers.SendErrorResponse(c, errno.New(errno.DatabaseError, err))
				return
			}
		}
	}

	// 签发 token
	tokenInfo, e := token.Sign(user)
	if e != nil {
		controllers.SendErrorResponse(c, e)
		return
	}

	uvw := viewmodels.NewUserAPISerializer(user)
	controllers.SendOKResponse(c, utils.MergeMap(uvw, map[string]interface{}{
		"meta": tokenInfo,
	}))
}

// Store 登录 (获取 token)
func Store(c *gin.Context) {
	var req *authorizationRequest.Login
	if err := c.ShouldBind(&req); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}

	user, err := req.Run()
	if err != nil {
		controllers.SendErrorResponse(c, err)
		return
	}

	// 签发 token
	tokenInfo, e := token.Sign(user)
	if e != nil {
		controllers.SendErrorResponse(c, e)
		return
	}

	uvw := viewmodels.NewUserAPISerializer(user)
	controllers.SendOKResponse(c, utils.MergeMap(uvw, map[string]interface{}{
		"meta": tokenInfo,
	}))
}

// Update 刷新 token
func Update(c *gin.Context, currentUser *userModel.User, tokenString string) {
	tokenInfo, err := token.Refresh(tokenString)
	if err != nil {
		controllers.SendErrorResponse(c, err)
		return
	}

	controllers.SendOKResponse(c, tokenInfo)
}

// Destroy 删除 token
func Destroy(c *gin.Context, currentUser *userModel.User, tokenString string) {
	token.Forget(tokenString, time.Duration(0))
	controllers.SendOKResponse(c, nil)
}
