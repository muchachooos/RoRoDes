package main

import (
	"GameAPI/configuration"
	"github.com/gin-gonic/gin"
	"strconv"
)

const configPath = "./configuration.json"

func main() {
	config := configuration.GetConfig(configPath)

	engine := gin.Default()

	engine.POST("/some_path", func(context *gin.Context) {})

	err := engine.Run(":" + strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}
}
