package model

import (
	. "BlogApi/utils"
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func New() {
	var err error
	// 读取数据库配置
	conf, err := GetMysqlConf("test", "")
	if err != nil {
		Logger.Errorf("读取mysql配置错误：\t", err)
	}
	// 连接数据库
	DB, err = gorm.Open(mysql.New(mysql.Config{
		// DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DB, conf.Charset),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "t_",                              // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true,                              // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	})
	if err != nil {
		Logger.Errorf("connect to %v error：%v", conf.Host, err)
	}
	DB.AutoMigrate(&Article{}, &Tag{}, &Category{}, &User{}, &UserRefArticle{}, &ArticleRefCategory{}, &ArticleRefTag{})
}
