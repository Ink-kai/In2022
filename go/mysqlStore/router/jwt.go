package router

import (
	"net/http"
	utils "pro22/mysqlStore/utils"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) {
	var (
		message gin.H
		token   string
		err     error
	)
	conf := utils.GetServerConf()
	if token, err = utils.GenerateToken(conf.Servername, conf.Serverpwd); err != nil {
		message = gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Token生成失败",
			"err":     err,
		}
	}
	message = gin.H{
		"code":  http.StatusInternalServerError,
		"token": token,
	}
	c.PureJSON(http.StatusOK, message)
}
