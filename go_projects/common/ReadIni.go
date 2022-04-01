package common

import (
	"gopkg.in/ini.v1"
)

type inkIni struct {
	path string
	data map[string]string
	// ini文件句柄
	f_handle *ini.File
	// ini文件分区句柄
	section_handle *ini.Section
}
type InIer interface {
	// 获取ini文件所有map
	InI_GetAllData() map[string]string
	// ini参数校验
	Validate([]string)
}

func New(path string) InIer {
	return &inkIni{
		path: path,
	}
}
func (i *inkIni) InI_GetAllData() map[string]string {
	f, _ := ini.LoadSources(ini.LoadOptions{SkipUnrecognizableLines: true}, i.path)
	default_sec, err := f.GetSection(ini.DefaultSection)
	if err != nil {
		panic(err)
	}
	AllMaps := readIniValue(default_sec)
	i.f_handle = f
	i.section_handle = default_sec
	i.data = AllMaps
	return i.data
}
func (i *inkIni) Validate(cols []string) {
	var errs string
	for _, v := range cols {
		result := i.section_handle.HasKey(v)
		if result == false {
			errs += v + " "
		}
	}
	if len(errs) > 0 {
		panic("请添加ini文件参数：" + errs)
	}
}

func readIniValue(section *ini.Section) map[string]string {
	var keysMap = make(map[string]string)
	for _, v := range section.KeyStrings() {
		keysMap[v] = section.Key(v).String()
	}
	return keysMap
}
