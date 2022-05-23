package common

import (
	"github.com/wonderivan/logger"

	"gopkg.in/ini.v1"
)

func Ini_loadSources(filepath, section string, options ini.LoadOptions, other ...interface{}) *ini.Section {
	cfg, err := ini.LoadSources(options, filepath)
	if err != nil {
		logger.Fatal("Fail to read file: %v", err)
	}
	var def_sec *ini.Section
	if section == "" {
		def_sec = cfg.Section(ini.DefaultSection)
	}
	def_sec = cfg.Section(section)
	return def_sec
}

func Ini_load(filepath, section string, other ...interface{}) *ini.Section {
	cfg, err := ini.Load(filepath)
	if err != nil {
		logger.Fatal("Fail to read file: %v", err)
	}
	var def_sec *ini.Section
	if section == "" {
		def_sec = cfg.Section(ini.DefaultSection)
	}
	def_sec = cfg.Section(section)
	return def_sec
}
