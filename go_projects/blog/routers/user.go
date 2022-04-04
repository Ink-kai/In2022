package routers

import (
	"blog/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
)

type user struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Phone    string `form:"phone" json:"phone"`
	State    int    `form:"state" json:"state" default:1`
	Address  string `form:"address" json:"address"`
}

func AddUser(c *gin.Context) {
	var u user
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service := services.User{
		Name:        u.Name,
		Password:    u.Password,
		Phone:       u.Phone,
		Address:     u.Address,
		CreateTime:  carbon.Now().Carbon2Time(),
		UpdatedTime: carbon.Now().Carbon2Time(),
	}
	if err := service.Add(); err != nil {
		c.JSON(http.StatusOK, u)
	}
	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
