package routers

import (
	"blog/models"
	"blog/pkg/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Account     int    `form:"account" json:"account"`
	Name        string `form:"name" json:"name" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	Phone       string `form:"phone" json:"phone"`
	State       int    `form:"state" json:"state" default:1`
	Address     string `form:"address" json:"address"`
	Email       string `form:"email" json:"email"`
	Createduser string `form:"createduser" json:"createduser"`
	Updateduser string `form:"updateduser" json:"updateduser"`
}

func AddUser(c *gin.Context) {
	var u user
	if err := c.ShouldBind(&u); err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model_user := models.UserNew(models.UserInfo{
		Uid:      common.Uuid_generate_v3(),
		Account:  u.Account,
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		Phone:    u.Phone,
		Address:  u.Address,
	})
	err := model_user.CreateUser()
	if err != nil {
		response = gin.H{
			"message": "success",
			"code":    http.StatusOK,
		}
		c.SecureJSON(http.StatusOK, response)
	} else {
		response = gin.H{
			"message": "参数错误",
			"code":    http.StatusInternalServerError,
		}
		c.SecureJSON(http.StatusOK, response)
	}
}
func DelUser(c *gin.Context) {
	err := models.DeleteUser(c.Param("uid"))
	if err != nil {
		response = gin.H{
			"message": "success",
			"code":    http.StatusOK,
		}
		c.SecureJSON(http.StatusOK, response)
	} else {
		response = gin.H{
			"message": "参数错误",
			"code":    http.StatusInternalServerError,
		}
		c.SecureJSON(http.StatusOK, response)
	}
}
func UpdateUser(c *gin.Context) {
	var u user
	if err := c.ShouldBind(&u); err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model_user := models.UserNew(models.UserInfo{
		Account:  u.Account,
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		Phone:    u.Phone,
		Address:  u.Address,
		// BasicModel: models.BasicModel{
		// 	UpdatedUser: u.Updateduser,
		// },
	})
	err := model_user.UpdateUser(c.Param("uid"))
	if err != nil {
		response = gin.H{
			"message": "success",
			"code":    http.StatusOK,
		}
		c.SecureJSON(http.StatusOK, response)
	} else {
		response = gin.H{
			"message": "参数错误",
			"code":    http.StatusInternalServerError,
		}
		c.SecureJSON(http.StatusOK, response)
	}
}
func GetAllUser(c *gin.Context) {
	users, err := models.GetAllUser()
	if err != nil {
		response = gin.H{
			"message": "success",
			"code":    http.StatusOK,
			"users":   users,
		}
		c.SecureJSON(http.StatusOK, response)
	} else {
		response = gin.H{
			"message": "参数错误",
			"code":    http.StatusInternalServerError,
			"users":   users,
		}
		c.SecureJSON(http.StatusOK, response)
	}
}
func GetUser(c *gin.Context) {
	user, err := models.GetUser(c.Param("uid"))
	if err != nil {
		response = gin.H{
			"message": "success",
			"code":    http.StatusOK,
			"user":    user,
		}
		c.SecureJSON(http.StatusOK, response)
	} else {
		response = gin.H{
			"message": "参数错误",
			"code":    http.StatusInternalServerError,
			"user":    user,
		}
		c.SecureJSON(http.StatusOK, response)
	}
}
func Login(c *gin.Context) {
	var u user
	if err := c.ShouldBind(&u); err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
