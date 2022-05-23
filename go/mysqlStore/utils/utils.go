package utils

import (
	"pro22/mysqlStore/common"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/wonderivan/logger"
	"gopkg.in/ini.v1"
)

type mysql_conf struct {
	IP       string
	Username string
	Password string
	Port     int
	DB       string
	Charset  string
	Timeout  int
}
type server_conf struct {
	IP           string
	Port         int
	Servername   string
	Serverpwd    string
	Secret       string
	Timeout      int64
	Tokentimeout int64
}
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var (
	nowtime      = time.Now()
	TOKENtimeout = time.Duration(GetServerConf().Tokentimeout)
	jwtsecret    = []byte(GetServerConf().Secret)
	ExpireTime   = nowtime.Add(TOKENtimeout * time.Second).Unix()
)

// 生成Token
func GenerateToken(username, password string) (string, error) {
	secret := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: ExpireTime,
			Issuer:    username,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS384, secret)
	token, err := tokenClaims.SignedString(jwtsecret)
	return token, err
}

// 验证Token
func VerifyToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtsecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// server
func GetServerConf() *server_conf {
	filepath := "H:\\学习资料\\In2022\\go\\mysqlStore\\conf\\server.ini"
	sec_name := "debug"
	section := common.Ini_loadSources(filepath, sec_name, ini.LoadOptions{
		SkipUnrecognizableLines: true,
	})
	port, err := section.Key("port").Int()
	if err != nil {
		logger.Fatal("读取server配置出现port错误", err)
	}
	timeout, err := section.Key("timeout").Int64()
	if err != nil {
		logger.Fatal("读取server配置出现timeout错误", err)
	}
	tokentimeout, err := section.Key("tokentimeout").Int64()
	if err != nil {
		logger.Fatal("读取server配置出现tokentimeout错误", err)
	}
	return &server_conf{
		IP:           section.Key("ip").String(),
		Port:         port,
		Servername:   section.Key("name").String(),
		Serverpwd:    section.Key("password").String(),
		Secret:       section.Key("secret").String(),
		Timeout:      timeout,
		Tokentimeout: tokentimeout,
	}
}

// mysql
func GetMysqlConf() *mysql_conf {
	filepath := "H:\\学习资料\\In2022\\go\\mysqlStore\\conf\\mysql.ini"
	sec_name := "debug"
	debug := common.Ini_loadSources(filepath, sec_name, ini.LoadOptions{
		SkipUnrecognizableLines: true,
	})
	port, _ := debug.Key("port").Int()
	timeout, err := debug.Key("timeout").Int()
	if err != nil {
		logger.Debug("my.ini配置文件int类型错误", err)
	}
	return &mysql_conf{
		IP:       debug.Key("ip").String(),
		Username: debug.Key("username").String(),
		Password: debug.Key("password").String(),
		DB:       debug.Key("db").String(),
		Port:     port,
		Charset:  debug.Key("charset").String(),
		Timeout:  timeout,
	}
}
