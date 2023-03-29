package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"textProject/src/common/database"
	"textProject/src/common/route"

	"github.com/gin-gonic/gin"
)

func main() {
	InitConfig()
	database.InitDB()
	ginService := gin.New()
	ginService = route.PathRoute(ginService)
	ginService.Run()
}

func InitConfig() {
	var envPara = "/src/config"
	workDir, _ := os.Getwd()
	viper.SetConfigName("base")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + envPara)
	fmt.Println(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
