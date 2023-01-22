package handler

import (
	"GameAPI/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Server) CreateUnitHandler(context *gin.Context) {
	cardID, ok := context.GetQuery("card_id")
	if cardID == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Card ID is missing"})
		return
	}

	gameID, ok := context.GetQuery("game_id")
	if gameID == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Game ID is missing"})
		return
	}

	fieldNumInString, ok := context.GetQuery("field_num")
	if fieldNumInString == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Field number is missing"})
		return
	}

	fieldNum, err := strconv.Atoi(fieldNumInString)

	unit, err := s.Storage.CreateUnitInDB(cardID, gameID, fieldNum)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	if len(unit) == 0 {
		context.JSON(http.StatusNotFound, model.Err{Error: "No card with this ID: " + err.Error()})
	}

	context.JSON(http.StatusOK, unit)
}
