package models

import (
	"blog/pkg/setting"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db *gorm.DB
)

// gorm.Model 的定义
type BasicModel struct {
	Updated     time.Time `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	Created     time.Time `gorm:"autoCreateTime"`       // 使用时间戳秒数填充创建时间
	CreatedUser string    `gorm:"not null"`
	UpdatedUser string    `gorm:"not null"`
}

func init() {
	var err error
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", setting.UserName, setting.Password, setting.IP, setting.Port, setting.DB)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // 日志级别
			Colorful:      false,         // 禁用彩色打印
		},
	)

	db, err = gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			// NoLowerCase:   true, // skip the snake_casing of names
		},
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
}
