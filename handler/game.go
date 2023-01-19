package handler

import (
	"GameAPI/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (d *Server) InitGameHandler(context *gin.Context) {
	gameID, err := d.Storage.InitGameInDB()
	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, gameID)
}

func (d *Server) GetGameHandler(context *gin.Context) {
	gameID, ok := context.GetQuery("id")
	if gameID == "" || !ok {
		context.Writer.WriteString("Game ID is missing")
		return
	}

	var game []model.Field

	err := d.Storage.DB.Select(&game, "SELECT `number`,`unit_id` FROM fields WHERE `game_id` = ?", gameID)
	if err != nil {
		panic(err)
	}

	if len(game) == 0 {
		context.Status(404)
		context.Writer.WriteString("No game with this ID")
		return
	}

	context.JSON(http.StatusOK, game)
}
