package setting

import (
	"time"

	"gopkg.in/ini.v1"
)

var (
	Addr         string
	Mode         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	Engine   string
	IP       string
	UserName string
	Password string
	Port     int
	DB       string
)

func init() {
	i, err := ini.LoadSources(ini.LoadOptions{SkipUnrecognizableLines: true}, "H:\\学习资料\\In2022\\go_projects\\blog\\conf\\blog.ini")
	if err != nil {
		panic(err)
	}
	Mode = GetKey(i, "", "mode").MustString("debug")
	// server
	ip := GetKey(i, "server", "ip").String()
	http_port := GetKey(i, "server", "http_port").String()
	Addr = ip + ":" + http_port
	ReadTimeout = time.Duration(GetKey(i, "server", "READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(GetKey(i, "server", "WRITE_TIMEOUT").MustInt(60)) * time.Second
	// mysql
	UserName = GetKey(i, "mysql", "name").String()
	Password = GetKey(i, "mysql", "password").String()
	Port = GetKey(i, "mysql", "port").MustInt(1433)
	IP = GetKey(i, "mysql", "ip").String()
	DB = GetKey(i, "mysql", "db").String()
	e := GetKey(i, "mysql", "engine")
	if e != nil {
		Engine = e.MustString("mysql")
	} else {
		Engine = "mysql"
	}
}

func GetKey(i *ini.File, section, key string) *ini.Key {
	var val *ini.Key
	if i.Section(section).HasKey(key) == true {
		val = i.Section(section).Key(key)
	} else {
		val = nil
	}
	return val
}
