package config

import (
	"github.com/spf13/viper"
)

// 应用程序配置
type appConfig struct {
	// 应用名称
	Name string
	// 运行模式: debug, release, test
	RunMode string
	// 运行 addr
	Addr string
	// 完整 url
	URL string
	// secret key
	Key string

	// 静态资源存放路径
	PublicPath string
	// 模板等前端源码文件存放路径
	ResourcesPath string
	// 模板文件存放的路径
	ViewsPath string

	// 是否开启 csrf
	EnableCsrf bool
	// csrf param name
	CsrfParamName string
	// csrf header
	CsrfHeaderName string

	// auth session key
	AuthSessionKey string
	// Context 中当前用户数据的 key
	ContextCurrentUserDataKey string

	// 百度翻译
	BaiduTranslateAppID string
	BaiduTranslateKey   string

	// 云片
	YunPianAPIKey string

	// 微信
	WeixinAppID       string
	WeixinAppSecret   string
	WeixinRedirectURL string

	// 极光推送
	JPushKey    string
	JPushSecret string
}

func newAppConfig() *appConfig {
	// 默认配置
	viper.SetDefault("APP.NAME", "gin_bbs")
	viper.SetDefault("APP.RUNMODE", "release")
	viper.SetDefault("APP.ADDR", ":8080")
	viper.SetDefault("APP.KEY", "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5")
	viper.SetDefault("APP.ENABLE_CSRF", true)

	return &appConfig{
		Name:    viper.GetString("APP.NAME"),
		RunMode: viper.GetString("APP.RUNMODE"),
		Addr:    viper.GetString("APP.ADDR"),
		URL:     viper.GetString("APP.URL"),
		Key:     viper.GetString("APP.KEY"),

		PublicPath:    "public",
		ResourcesPath: "resources",
		ViewsPath:     "resources/views",

		EnableCsrf:     viper.GetBool("APP.ENABLE_CSRF"),
		CsrfParamName:  "_csrf",
		CsrfHeaderName: "X-CsrfToken",

		AuthSessionKey:            "gin_session",
		ContextCurrentUserDataKey: "currentUserData",

		BaiduTranslateAppID: viper.GetString("BAIDU_TRANSLATE.APPID"),
		BaiduTranslateKey:   viper.GetString("BAIDU_TRANSLATE.KEY"),

		YunPianAPIKey: viper.GetString("YUNPIAN_API_KEY"),

		WeixinAppID:       viper.GetString("WEIXIN.APP_ID"),
		WeixinAppSecret:   viper.GetString("WEIXIN.APP_SECRET"),
		WeixinRedirectURL: viper.GetString("WEIXIN.REDIRECT_URL"),

		JPushKey:    viper.GetString("JPUSH.KEY"),
		JPushSecret: viper.GetString("JPUSH.SECRET"),
	}
}
