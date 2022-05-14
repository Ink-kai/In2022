package v1

import (
	"bufio"
	"io"
	"net/http"
	service "pro22/mysqlStore/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

func UplodFiles(c *gin.Context) {
	var message string
	form, err := c.MultipartForm()
	if err != nil {
		c.PureJSON(http.StatusInternalServerError, gin.H{
			"message": "上传文件获取异常",
		})
	}
	files := form.File["files"]
	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			message = "文件格式错误"
		}
		bytesize := make([]byte, file.Size)
		reader := bufio.NewReader(f)
		for {
			_, err = reader.Read(bytesize)
			if err != nil || err != io.EOF {
				break
			}
		}
		_, remoteBool := c.RemoteIP()
		finfo := map[string]interface{}{
			"name":       file.Filename,
			"content":    bytesize,
			"size":       file.Size,
			"type":       strings.Split(file.Header.Get("Content-Type"), "/")[1],
			"source":     "纯接口调用",
			"createuser": "ink",
			"clientIp":   c.ClientIP(),
			"remoteBool": remoteBool,
		}
		// 通过mysql将数据写入
		err = service.Insert(finfo)
		if err != nil {
			message = "数据上传失败"
			logger.Fatal("数据上传失败。%v", err)
		}
		message = "文件上传成功"
	}
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func Example(c *gin.Context) {
	remote_ip, err := c.RemoteIP()
	logger.Info(err)
	msg := gin.H{
		"client": c.ClientIP(),
		"remote": remote_ip,
	}
	c.JSON(http.StatusOK, msg)
}
