package midware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//声明一个中间件方法
func MiddleW() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前")
		c.Next()
		fmt.Println("我在方法后")
	}
}
