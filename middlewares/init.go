package middlewares

// test 3

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init(c *gin.Context) {
	fmt.Println("中间件测试")
	c.Set("user", "LTXXXXXX!")
}
