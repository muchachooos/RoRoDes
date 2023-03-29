package handler

import (
	"GameAPI/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) InitGameHandler(context *gin.Context) {
	gameID, err := s.Storage.InitGameInDB()
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, gameID)
}

func (s *Server) GetGameHandler(context *gin.Context) {
	gameID, ok := context.GetQuery("game_id")
	if gameID == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Game ID is missing"})
		return
	}

	game, err := s.Storage.GetGameFromDB(gameID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	if len(game) == 0 {
		context.JSON(http.StatusNotFound, model.Err{Error: "No game with this ID: " + err.Error()})
	}

	context.JSON(http.StatusOK, game)
}
