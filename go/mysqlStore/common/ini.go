package common

import (
	"github.com/wonderivan/logger"

	"gopkg.in/ini.v1"
)

type mysql_conf struct {
	IP       string
	Username string
	Password string
	Port     int
	DB       string
	Charset  string
	Timeout  int
}

func GetMysqlConf() *mysql_conf {
	cfg, err := ini.LoadSources(ini.LoadOptions{
		SkipUnrecognizableLines: true,
	}, "H:\\学习资料\\In2022\\go\\mysqlStore\\confs\\mysql.ini")
	if err != nil {
		logger.Fatal("Fail to read file: %v", err)
	}
	debug := cfg.Section("debug")
	port, _ := debug.Key("port").Int()
	timeout, err := debug.Key("timeout").Int()
	if err != nil {
		logger.Debug("my.ini配置文件键值类型错误")
	}
	return &mysql_conf{
		IP:       debug.Key("ip").String(),
		Username: debug.Key("username").String(),
		Password: debug.Key("password").String(),
		DB:       debug.Key("db").String(),
		Port:     port,
		Charset:  debug.Key("charset").String(),
		Timeout:  timeout,
	}
}
