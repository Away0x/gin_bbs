package ginutils

import (
	"github.com/gin-gonic/gin"
)

type ConfigOption struct {
	URL         string
	PublicPath  string
	MixFilePath string // laravel-mix manifest.json 文件地址

	EnableCsrf       bool
	CsrfParamName    string
	CsrfHeaderName   string
	CsrfErrorHandler func(*gin.Context, bool) // 二参表示 csrf token 是从 header 中拿到的
}

var (
	config *ConfigOption
)

// InitGinUtils 初始化 ginutils 的配置
func InitGinUtils(options ConfigOption) {
	config = &ConfigOption{
		URL:              options.URL,
		PublicPath:       options.PublicPath,
		MixFilePath:      options.MixFilePath,
		EnableCsrf:       options.EnableCsrf,
		CsrfParamName:    options.CsrfParamName,
		CsrfHeaderName:   options.CsrfHeaderName,
		CsrfErrorHandler: options.CsrfErrorHandler,
	}

	if config.EnableCsrf && config.CsrfErrorHandler == nil {
		config.CsrfErrorHandler = func(c *gin.Context, inHeader bool) {}
	}
}

// 获取 ginutils 配置
func GetGinUtilsConfig() *ConfigOption {
	if config == nil {
		panic("[ginutils] config init error")
	}

	return config
}
