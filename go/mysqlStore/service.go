package main

import (
	_ "pro22/mysqlStore/models"
	_ "pro22/mysqlStore/service/server"
	_ "pro22/mysqlStore/utils"

	"github.com/wonderivan/logger"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			logger.Fatal(err)
		}
	}()
}
