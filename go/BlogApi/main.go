package main

import (
	service "BlogApi/service"
	. "BlogApi/utils"
	"fmt"
)

func main() {
	// model.New()
	conf, err := GetServerConf("test", "conf/server.ini")
	if err != nil {
		Log.Errorf("读取server配置错误：\t%+v", err)
	}
	router := service.Setup()
	router.Run(fmt.Sprintf("%s:%d", conf.Host, conf.Port))
}
