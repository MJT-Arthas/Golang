package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// 简单的路由组: v1
	v1 := router.Group("/gapi")
	{
		v1.GET("/", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "hello world",
			})
		})
		v1.GET("/test", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "test",
			})
		})

	}

	router.Run(":3001")
}
