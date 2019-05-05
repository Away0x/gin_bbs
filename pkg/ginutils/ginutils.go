package ginutils

import (
	"path"
)

// ConfigOption -
type ConfigOption struct {
	URL         string
	PublicPath  string
	MixFilePath string // laravel-mix manifest.json 文件地址

	EnableCsrf     bool
	CsrfParamName  string
	CsrfHeaderName string
}

var config *ConfigOption

// InitGinUtils 初始化 ginutils 的配置
func InitGinUtils(options ConfigOption) {
	config = &ConfigOption{
		URL:            options.URL,
		PublicPath:     options.PublicPath,
		MixFilePath:    path.Join(options.PublicPath, "mix-manifest.json"),
		EnableCsrf:     options.EnableCsrf,
		CsrfParamName:  options.CsrfParamName,
		CsrfHeaderName: options.CsrfHeaderName,
	}

	if options.MixFilePath != "" {
		config.MixFilePath = options.MixFilePath
	}
}

// GetGinUtilsConfig 获取 ginutils 配置
func GetGinUtilsConfig() *ConfigOption {
	if config == nil {
		panic("[ginutils] config init error")
	}

	return config
}
