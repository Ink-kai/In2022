package middleware

import (
	"net/http"
	utils "pro22/mysqlStore/utils"
	status "pro22/mysqlStore/utils/code"
	"time"

	"github.com/wonderivan/logger"

	"github.com/gin-gonic/gin"
)

type Message struct {
	ClientIP string
	Path     string
	Code     int
	Message  string
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			message Message
		)
		token := c.GetHeader("Token")
		message = Message{
			c.ClientIP(),
			c.FullPath(),
			http.StatusOK,
			token,
		}
		if token == "" {
			message = Message{
				c.ClientIP(),
				c.FullPath(),
				status.INVALID_PARAMS,
				"无效参数",
			}
			logger.Debug(message)
		} else {
			claims, err := utils.VerifyToken(token)
			if err != nil {
				message = Message{
					c.ClientIP(),
					c.FullPath(),
					status.ERROR_AUTH_CHECK_TOKEN_FAIL,
					"Token失效",
				}
			} else if time.Now().Unix() > claims.ExpiresAt {
				message = Message{
					c.ClientIP(),
					c.FullPath(),
					status.ERROR_AUTH_CHECK_TOKEN_TIMEOUT,
					"Token超时",
				}
			}
			logger.Debug(message)
		}
		if message.Code != http.StatusOK {
			message = Message{
				c.ClientIP(),
				c.FullPath(),
				http.StatusUnauthorized,
				"Full authentication is required to access this resource.",
			}
			logger.Debug(message)
			c.PureJSON(message.Code, message)
			c.Abort()
			return
		}
		c.Next()
	}
}
