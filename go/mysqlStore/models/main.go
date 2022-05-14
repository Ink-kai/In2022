package model

import (
	"fmt"
	"log"
	"os"
	common "pro22/mysqlStore/common"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db    *gorm.DB
	dbErr error
)

func IninDB() {
	conf := common.GetMysqlConf()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local&timeout=%vs", conf.Username, conf.Password, conf.IP, conf.Port, conf.DB, conf.Charset, conf.Timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, dbErr = gorm.Open("mysql", dsn)
	if dbErr != nil {
		panic("连接数据库失败, error=" + dbErr.Error())
	}
	// db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	// 禁用复数表名
	db.SingularTable(true)
	// 自动生成
	db.AutoMigrate(&FileInfo{})
}
