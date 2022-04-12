package routers

import (
	. "blog/pkg/setting"

	"github.com/gin-gonic/gin"
)

var (
	response = make(map[string]interface{})
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
		api.GET("/user/:uid", GetUser)
		api.GET("/users", GetAllUser)
		api.GET("/comment")

		// 修改
		api.PUT("/article:aid")
		api.PUT("/tags:tid")
		api.PUT("/user/:uid", UpdateUser)
		api.PUT("/comment:cid")

		// 删除
		api.DELETE("/article:aid")
		api.DELETE("/tags:tid")
		api.DELETE("/user/:uid", DelUser)
		api.DELETE("/comment:cid")
	}
	return r
}
