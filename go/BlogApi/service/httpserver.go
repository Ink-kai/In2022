package service

import (
	"BlogApi/model"
	. "BlogApi/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() {
	model.New()
	conf, err := GetServerConf("test")
	if err != nil {
		Logger.Errorf("读取server配置错误：\t", err)
	}
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "welcome to my site",
		})
	})
	router.Run(fmt.Sprintf("%s:%d", conf.Host, conf.Port))
}
