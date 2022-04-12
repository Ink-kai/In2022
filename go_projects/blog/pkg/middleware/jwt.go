package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthToken(c *gin.Context) {
	t := time.Now()
	fmt.Fprintf(os.Stdout, "%v\n", "路由中间件开始....")
	c.Next()
	fmt.Fprintf(os.Stdout, "%v\t%v\n", "路由中间件结束。", time.Since(t))
}
