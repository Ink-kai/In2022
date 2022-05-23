package utils

import "github.com/wonderivan/logger"

func InitLogger() {
	filepath := ""
	// 通过配置文件配置
	logger.SetLogger(filepath)
}
