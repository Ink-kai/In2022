package router

import (
	"bufio"
	"io"
	"mime/multipart"
	"net/http"
	service "pro22/mysqlStore/service/file"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var (
	wg sync.WaitGroup
)

func UplodFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.PureJSON(http.StatusInternalServerError, gin.H{
			"message": "上传文件获取异常",
		})
	}
	files := form.File["files"]
	allMessage := make(chan string)
	for _, file := range files {
		wg.Add(1)
		go DisposeFile(file, c, allMessage)
		wg.Done()
	}
	wg.Wait()
	c.PureJSON(http.StatusOK, gin.H{
		"message": <-allMessage,
	})
}
func DisposeFile(file *multipart.FileHeader, c *gin.Context, message chan string) {
	f, err := file.Open()
	if err != nil {
		message <- "文件格式错误"
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
		"type":       file.Header.Get("Content-Type"),
		"source":     "纯接口调用",
		"createuser": "ink",
		"clientIp":   c.ClientIP(),
		"remoteBool": remoteBool,
	}
	err = service.FILE_Insert(finfo)
	if err != nil {
		message <- "数据上传失败"
		logger.Fatal("数据上传失败。%v", err)
	}
	message <- "文件上传成功"
}
