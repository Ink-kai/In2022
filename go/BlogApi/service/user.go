package service

import (
	. "BlogApi/model"
	. "BlogApi/utils"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserForm struct {
	Name     string `form:"name"	json:"name"	validate:"required"`
	Password string `form:"password"	json:"password"	validate:"required"`
	Email    string `form:"email" json:"email" validate:"omitempty,email"`
	Phone    string `form:"phone"	json:"phone"`
	Birthday string `form:"birthday"	json:"birthday"`
	Address  string `form:"address"	json:"address"`
	Desc     string `form:"desc"	json:"desc"`
	Remark   string `form:"remark"	json:"remark"`
}

func AddUser(ctx *gin.Context) {
	var (
		data     UserForm
		validate       = validator.New()
		msg            = gin.H{}
		Birthday int64 = time.Now().Unix()
	)
	if err := ctx.Bind(&data); err != nil {
		Log.Errorf("绑定FormUser结构体错误。\t%+v", err.Error())
	}
	body, _ := io.ReadAll(ctx.Request.Body)
	if len(body) != 0 {
		json.Unmarshal(body, &(data))
	} else if data.Name == "" && data.Password == "" {
		msg = gin.H{
			"code":    http.StatusInternalServerError,
			"message": "参数不完整。请查看参数校验",
		}
		ctx.PureJSON(http.StatusOK, msg)
		return
	}
	UserUID := GenerateID()
	data.Password = string(EncryptPassword(data.Password))
	if data.Birthday != "" {
		Birthday = TimeStrToTimeDefault(data.Birthday).Unix()
	}
	err := validate.Struct(data)
	if err != nil {
		var validate_err string
		for _, err := range err.(validator.ValidationErrors) {
			validate_err = err.Error()
			Log.Errorf("%v	", err)
		}
		msg = gin.H{
			"code":    http.StatusInternalServerError,
			"message": validate_err,
		}
		ctx.PureJSON(http.StatusOK, msg)
		return
	}
	cookie, _ := ctx.Cookie("mode")
	if cookie != "debug" {
		u := &User{
			UserUID:  UserUID,
			Name:     data.Name,
			Password: data.Password,
			Email:    data.Email,
			Phone:    data.Phone,
			Birthday: Birthday,
			Address:  data.Address,
			Remark:   data.Remark,
			Desc:     data.Desc,
		}
		if err := u.AddUser(); err != nil {
			Log.Errorf("创建用户错误：\t%+v", err)
		}
	}
	msg = gin.H{
		"code": http.StatusOK,
		"message": gin.H{
			"name":  data.Name,
			"email": data.Email,
		},
	}
	ctx.PureJSON(http.StatusOK, msg)
	return
}
