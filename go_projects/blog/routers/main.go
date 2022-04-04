package routers

import (
	. "blog/pkg/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(Mode)
	gin.DisableConsoleColor()
	api := r.Group("api")
	{
		// 新增
		api.POST("/user", AddUser)
		api.POST("/article")
		api.POST("/tags")
		api.POST("/comment")

		// 获取
		api.GET("/article")
		api.GET("/tags")
		api.GET("/user")
		api.GET("/comment")

		// 修改
		api.PUT("/article:aid")
		api.PUT("/tags:aid")
		api.PUT("/user:aid")
		api.PUT("/comment:aid")

		// 删除
		api.DELETE("/article:aid")
		api.DELETE("/tags:aid")
		api.DELETE("/user:aid")
		api.DELETE("/comment:aid")
	}
	return r
}
