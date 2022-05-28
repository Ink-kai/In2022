package utils

import (
	go_logger "github.com/phachon/go-logger"
)

var Logger *go_logger.Logger

func init() {
	Logger = go_logger.NewLogger()
	Logger.Detach("console")

	// 命令行输出配置
	consoleConfig := &go_logger.ConsoleConfig{
		Color:      true, // 命令行输出字符串是否显示颜色
		JsonFormat: true, // 命令行输出字符串是否格式化
		Format:     "",   // 如果输出的不是 json 字符串，JsonFormat: false, 自定义输出的格式
	}
	Logger.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)
}
