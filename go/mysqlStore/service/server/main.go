package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"pro22/mysqlStore/middleware"
	route "pro22/mysqlStore/router"
	utils "pro22/mysqlStore/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

func init() {
	conf := utils.GetServerConf()
	gin.SetMode("")
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", conf.IP, conf.Port),
		Handler:        router,
		ReadTimeout:    time.Duration(conf.Timeout) * time.Second,
		WriteTimeout:   time.Duration(conf.Timeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	root := router.Group("/")
	{
		root.GET("getToken", route.GetToken)
	}
	api := router.Group("/api")
	api.Use(middleware.JWT())
	{
		api.POST("uploadFile", route.UplodFiles)
	}
	go func() {
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
}
