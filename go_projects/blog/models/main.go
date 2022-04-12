package models

import (
	"blog/pkg/setting"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// gorm.Model 的定义
type BasicModel struct {
	ID          int `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedUser string
	UpdatedUser string
}

func init() {
	var err error
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", setting.UserName, setting.Password, setting.IP, setting.Port, setting.DB)
	db, err = gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n\n", 0))
	db.AutoMigrate(UserInfo{}, Article{})
}
