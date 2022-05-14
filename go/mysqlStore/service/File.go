package FileService

import (
	model "pro22/mysqlStore/models"

	"github.com/wonderivan/logger"
)

func Insert(data map[string]interface{}) error {
	file := model.FileNew()
	file.Name = data["name"].(string)
	file.Content = data["content"].([]byte)
	file.Size = data["size"].(int64)
	file.Filetype = data["type"].(string)
	file.Source = data["source"].(string)
	createuser := data["createuser"].(string)
	file.ClientIp = data["clientIp"].(string)
	file.RemoteBool = data["remoteBool"].(bool)
	if createuser == "" {
		file.CreateUser = "未知"
	} else {
		file.CreateUser = createuser
	}
	if err := file.Insert(); err != nil {
		logger.Fatal("%v 数据写入失败。%v", file.Name, err)
	}
	return nil
}
