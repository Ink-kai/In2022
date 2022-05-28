package utils

import (
	"gopkg.in/ini.v1"
)

type MysqlConf struct {
	Host     string
	User     string
	Password string
	Port     int
	DB       string
	Charset  string
}

type ServerConf struct {
	Host string
	Port int
}

func GetServerConf(mode, path string) (*ServerConf, error) {
	var (
		cfg *ini.File
		err error
		sec *ini.Section
	)
	if path == "" {
		cfg, err = ini.LoadSources(ini.LoadOptions{
			SkipUnrecognizableLines: true,
		}, "conf/server.ini")
	} else {
		cfg, err = ini.LoadSources(ini.LoadOptions{
			SkipUnrecognizableLines: true,
		}, path)
	}
	if err != nil {
		// Logger.Errorf("open ini file error：%v", err)
		return nil, err
	}

	if mode == "" {
		sec = cfg.Section(ini.DEFAULT_SECTION)
	} else {
		sec, err = cfg.GetSection(mode)
		if err != nil {
			// Logger.Errorf("%v section not exists：%v", mode, err)
			return nil, err
		}
	}
	conf := &ServerConf{
		Host: sec.Key("host").MustString("127.0.0.1"),
		Port: sec.Key("port").MustInt(8086),
	}
	return conf, nil
}

func GetMysqlConf(mode, path string) (*MysqlConf, error) {

	var (
		cfg *ini.File
		err error
		sec *ini.Section
	)
	if path == "" {
		cfg, err = ini.LoadSources(ini.LoadOptions{
			SkipUnrecognizableLines: true,
		}, "conf/mysql.ini")
	} else {
		cfg, err = ini.LoadSources(ini.LoadOptions{
			SkipUnrecognizableLines: true,
		}, path)
	}
	if err != nil {
		// Logger.Errorf("open ini file error：%v", err)
		return nil, err
	}

	if mode == "" {
		sec = cfg.Section(ini.DEFAULT_SECTION)
	} else {
		sec, err = cfg.GetSection(mode)
		if err != nil {
			return nil, err
			// Logger.Errorf("%v section not exists：%v", mode, err)
		}
	}
	conf := &MysqlConf{
		Host:     sec.Key("host").MustString("127.0.0.1"),
		User:     sec.Key("username").MustString("sa"),
		Password: sec.Key("password").MustString("sa"),
		Port:     sec.Key("port").MustInt(3306),
		DB:       sec.Key("db").MustString("master"),
		Charset:  sec.Key("charset").MustString("utf-8"),
	}
	return conf, nil
}
