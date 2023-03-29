package main

import (
	"GameAPI/configuration"
	"GameAPI/handler"
	"GameAPI/storage"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
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

	server := handler.Server{
		Storage: &storage.Storage{
			DB: dataBase,
		},
	}

	engine := gin.Default()

	engine.POST("/init_game", server.InitGameHandler)
	engine.GET("/get_game", server.GetGameHandler)
	engine.POST("/create_unit", server.CreateUnitHandler)
	engine.GET("/get_unit", server.GetUnitHandler)
	engine.POST("/move_unit", server.MoveUnitHandler)

	// Файловый сервер который возвращает html, css, js и другие файлы
	engine.StaticFS("/game", http.Dir("client"))

	// Перенаправляем все запросы без относительного пути, пример: "www.here.com" -> "www.here.com/client"
	engine.GET("", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "client")
	})

	err = engine.Run(":" + strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}
}
