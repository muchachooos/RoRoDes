package main

import (
	"GameAPI/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func main() {
	configInBytes, err := os.ReadFile("./configuration.json")
	if err != nil {
		panic(err)
	}

	var config model.Config

	err = json.Unmarshal(configInBytes, &config)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()

	engine.POST("/some_path", func(context *gin.Context) {})

	err = engine.Run(":" + strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}
}
