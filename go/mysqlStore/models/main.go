package model

import (
	"fmt"
	"log"
	"os"
	utils "pro22/mysqlStore/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db    *gorm.DB
	dbErr error
)

func init() {
	conf := utils.GetMysqlConf()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local&timeout=%vs", conf.Username, conf.Password, conf.IP, conf.Port, conf.DB, conf.Charset, conf.Timeout)
	db, dbErr = gorm.Open("mysql", dsn)
	if dbErr != nil {
		panic("连接数据库失败, error=" + dbErr.Error())
	}
	// 禁用复数表名
	db.SingularTable(true)
	db.AutoMigrate(&fileInfo{})
	// db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))
}
