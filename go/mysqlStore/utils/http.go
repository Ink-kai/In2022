package utils

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	model "pro22/mysqlStore/models"
	v1 "pro22/mysqlStore/routes/v1"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

func Setup() *gin.Engine {
	model.IninDB()
	// 禁用控制台颜色
	gin.DisableConsoleColor()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		logger.Info("%v %v", httpMethod, absolutePath)
	}
	// 创建记录日志的文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 创建一个不包含中间件的路由器
	router := gin.New()
	// 全局中间件
	// 使用 Logger 中间件
	router.Use(gin.Logger())
	// 使用 Recovery 中间件
	router.Use(gin.Recovery())
	srv := &http.Server{
		Addr:           "192.168.207.1:8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ink",
		})
	})
	e := router.Group("/v1")
	{
		e.POST("uploadFile", v1.UplodFiles)
		e.GET("example", v1.Example)
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	logger.Info("Server exiting")
	return router
}
