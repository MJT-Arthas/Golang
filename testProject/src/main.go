package main

import (
	"textProject/src/common/route"

	"github.com/gin-gonic/gin"
)

func main() {
	ginService := gin.New()
	ginService = route.PathRoute(ginService)
	ginService.Run()
}
