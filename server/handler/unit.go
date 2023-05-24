package handler

import (
	"RoRoDes/model"
	"RoRoDes/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	xInString, ok := context.GetQuery("x")
	if xInString == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "X is missing"})
		return
	}

	yInString, ok := context.GetQuery("y")
	if yInString == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Y is missing"})
		return
	}

	x, err := strconv.Atoi(xInString)

	y, err := strconv.Atoi(yInString)

	unit, err := s.Service.CreateUnit(cardID, gameID, x, y)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	if len(unit) == 0 {
		context.JSON(http.StatusNotFound, model.Err{Error: "No card with this ID: " + cardID})
		return
	}

	context.JSON(http.StatusOK, unit)
}

func (s *Server) GetUnitHandler(context *gin.Context) {
	unitID, ok := context.GetQuery("unit_id")
	if unitID == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Unit ID is missing"})
		return
	}

	unitField, err := s.Service.GetUnit(unitID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	if len(unitField) == 0 {
		context.JSON(http.StatusNotFound, model.Err{Error: "No unit with this ID: " + unitID})
		return
	}

	context.JSON(http.StatusOK, unitField)
}

func (s *Server) MoveUnitHandler(context *gin.Context) {
	uintId, ok := context.GetQuery("unit_id")
	if uintId == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Unit ID is missing"})
		return
	}

	dir, ok := context.GetQuery("direction")
	if dir == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Direction is missing"})
		return
	}

	ok, err := s.Service.MoveUnit(uintId, dir)
	if err != nil {
		if err == storage.ErrMoveUnit {
			context.JSON(http.StatusConflict, model.Err{Error: err.Error()})
			return
		}

		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, ok)
}
