package service

import (
	. "BlogApi/model"
	. "BlogApi/utils"
)

func AddUser(data map[string]interface{}) error {
	user := &User{
		Uid:      data["uid"].(uint),
		Name:     data["name"].(string),
		Password: data["password"].(string),
		Birthday: data["birthday"].(uint),
		Email:    data["email"].(string),
		Address:  data["address"].(string),
		Phone:    data["phone"].(string),
		Desc:     data["desc"].(string),
		Remark:   data["remark"].(string),
	}
	if err := user.AddUser(); err != nil {
		Logger.Errorf("创建用户错误：%v", err)
	}
	return nil
}
