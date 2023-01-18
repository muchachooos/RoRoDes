package main

import (
	"GameAPI/configuration"
	"GameAPI/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
)

const configPath = "./configuration.json"

func main() {
	config := configuration.GetConfig(configPath)

	dataBase, err := sqlx.Open("mysql", config.DataSourceName)
	if err != nil {
		panic(err)
	}

	if dataBase == nil {
		panic("Data base nil")
	}

	server := handler.DataBase{
		DB: dataBase,
	}

	engine := gin.Default()

	engine.POST("/init_game", server.InitGameHandler)
	engine.GET("/get_game", server.GetGameHandler)

	err = engine.Run(":" + strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}
}
