package main

import (
	"RoRoDes/configuration"
	"RoRoDes/handler"
	"RoRoDes/model"
	"RoRoDes/service"
	"RoRoDes/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
)

const configPath = "configuration.json"

func main() {
	config := configuration.GetConfig(configPath)

	dataBase, err := sqlx.Open("mysql", getDSN(config.DBConf))
	if err != nil {
		panic(err)
	}

	if dataBase == nil {
		panic("Data base nil")
	}

	server := handler.Server{
		Service: &service.Service{
			Storage: &storage.Storage{
				DB: dataBase,
			},
		},
	}

	engine := gin.Default()

	engine.POST("/init_game", server.InitGameHandler)
	engine.GET("/get_all_game_id", server.GetAllGameIdHandler)
	engine.GET("/get_game", server.GetGameHandler)
	engine.POST("/create_unit", server.CreateUnitHandler)
	engine.GET("/get_unit", server.GetUnitHandler)
	engine.POST("/move_unit", server.MoveUnitHandler)
	engine.GET("/get_card", server.GetCardHandler)
	engine.GET("/get_all_name", server.GetNameAllCardsHandler)
	engine.PUT("/add_card_in_deck", server.AddCardInDeckHandler)
	engine.GET("/get_deck", server.GetDeckHandler)

	err = engine.Run(":" + strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}
}

func getDSN(cfg model.DBConf) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.DBPort, cfg.DBName)
}
