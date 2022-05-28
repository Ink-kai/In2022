package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "welcome to my site"})
	})
	test := router.Group("/api/test/")
	{
		test.POST("/AddUser", AddUser)
	}
	return router
}
