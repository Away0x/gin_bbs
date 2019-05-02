package config

import (
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

func initLog() {
	// 默认配置
	viper.SetDefault("LOG.WRITERS", "file,stdout")
	viper.SetDefault("LOG.LOGGER_LEVEL", "DEBUG")
	viper.SetDefault("LOG.LOGGER_FILE", logFilePath)
	viper.SetDefault("LOG.LOG_FORMAT_TEXT", false)
	viper.SetDefault("LOG.ROLLING_POLICY", "size")
	viper.SetDefault("LOG.LOG_ROTATE_DATE", 1)
	viper.SetDefault("LOG.LOG_ROTATE_SIZE", 1)
	viper.SetDefault("LOG.LOG_BACKUP_COUNT", 7)

	logConfig := log.PassLagerCfg{
		// 输出位置，有两个可选项 —— file 和 stdout。选择 file 会将日志记录到 logger_file 指定的日志文件中，选择 stdout 会将日志输出到标准输出，当然也可以两者同时选择
		Writers: viper.GetString("LOG.WRITERS"),
		// 日志级别，DEBUG、INFO、WARN、ERROR、FATAL
		LoggerLevel: viper.GetString("LOG.LOGGER_LEVEL"),
		// 日志文件
		LoggerFile: viper.GetString("LOG.LOGGER_FILE"),
		// 日志的输出格式，JSON 或者 plaintext，true 会输出成 JSON 格式，false 会输出成非 JSON 格式
		LogFormatText: viper.GetBool("LOG.LOG_FORMAT_TEXT"),
		// rotate 依据，可选的有 daily 和 size。如果选 daily 则根据天进行转存，如果是 size 则根据大小进行转存
		RollingPolicy: viper.GetString("LOG.ROLLING_POLICY"),
		// rotate 转存时间，配 合rollingPolicy: daily 使用
		LogRotateDate: viper.GetInt("LOG.LOG_ROTATE_DATE"),
		// rotate 转存大小，配合 rollingPolicy: size 使用 (大于 1mb 会压缩为 zip)
		LogRotateSize: viper.GetInt("LOG.LOG_ROTATE_SIZE"),
		// 当日志文件达到转存标准时，log 系统会将该日志文件进行压缩备份，这里指定了备份文件的最大个数
		LogBackupCount: viper.GetInt("LOG.LOG_BACKUP_COUNT"),
	}

	log.InitWithConfig(&logConfig)
}
