package RoleManage

import (
	"flag"
	"log"
	"net/url"
	. "pro22/common"
	"strconv"
	"strings"
)

func Run() {
	var path string
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	flag.StringVar(&path, "f", "H:\\学习资料\\In2022\\go_projects\\example_file\\t.ini", "本地ini配置文件")
	flag.Parse()
	//1.准备数据
	i := New(path)
	ini_data := i.InI_GetAllData()
	i.Validate([]string{"url", "cookie", "ExcelPath", "rolename", "rolecode_name", "rolecode_num", "sysid", "rolename_num", "username_num", "useruid_num"})
	sysId := ini_data["sysid"]
	postUrl := ini_data["url"]
	cookie := ini_data["cookie"]
	ExcelPath := ini_data["ExcelPath"]
	roleName := ini_data["rolename"]
	roleCodeName := ini_data["rolecode_name"]
	roleNameNum, _ := strconv.Atoi(ini_data["rolename_num"])
	roleCodeNum, _ := strconv.Atoi(ini_data["rolecode_num"])
	userNameNum, _ := strconv.Atoi(ini_data["username_num"])
	userIdNum, _ := strconv.Atoi(ini_data["useruid_num"])
	txt := Read_excel(ExcelPath, roleName, roleCodeName, roleNameNum, roleCodeNum, userNameNum, userIdNum)
	txt = "roleId=" + sysId + "&userInfoList=" + url.QueryEscape(txt)
	// //2.封装数据
	var r_meta = req_meta{
		Url:          postUrl,
		Method:       "POST",
		Url_form:     strings.NewReader(txt),
		Content_Type: "application/x-www-form-urlencoded",
		Cookie:       cookie,
	}
	// //3.发送数据到接口
	Request_url(r_meta)
}
