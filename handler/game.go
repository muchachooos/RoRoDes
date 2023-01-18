package handler

import (
	"GameAPI/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
)

func (s *DataBase) InitGameHandler(context *gin.Context) {
	bodyInBytes, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Read body error: " + err.Error()})
		return
	}

	err = context.Request.Body.Close()
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Close body error: " + err.Error()})
		return
	}

	var game []model.Field

	err = json.Unmarshal(bodyInBytes, &game)
	if err != nil {
		if err != nil {
			context.JSON(http.StatusBadRequest, model.Err{Error: "Unmarshal request body error: " + err.Error()})
		}
	}

	gameId := uuid.NewString()

	_, err = s.DB.Exec("INSERT INTO game (`id`) VALUES (?)", gameId)
	if err != nil {
		panic(err)
	}

	for i := range game {
		_, err = s.DB.Exec("INSERT INTO fields (`number`, `unit_id`, `game_id`) VALUES (?,?,?)", game[i].Number, game[i].UnitID, gameId)
		if err != nil {
			panic(err)
		}
	}

	context.JSON(http.StatusOK, gameId)
}

func (s *DataBase) GetGameHandler(context *gin.Context) {
	gameID, ok := context.GetQuery("id")
	if gameID == "" || !ok {
		context.Writer.WriteString("Game ID is missing")
		return
	}

	var game []model.Field

	err := s.DB.Select(&game, "SELECT `number`,`unit_id` FROM fields WHERE `game_id` = ?", gameID)
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
