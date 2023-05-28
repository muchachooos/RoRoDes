package handler

import (
	"RoRoDes/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) GetCardHandler(context *gin.Context) {
	cardID, ok := context.GetQuery("card_id")
	if cardID == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Card ID is missing"})
		return
	}

	characteristics, err := s.Service.GetCard(cardID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	if len(characteristics) == 0 {
		context.JSON(http.StatusNotFound, model.Err{Error: "No card with this ID: " + cardID})
		return
	}

	context.JSON(http.StatusOK, characteristics)
}

func (s *Server) GetNameAllCardsHandler(context *gin.Context) {
	allName, err := s.Service.GetName()
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	if len(allName) == 0 {
		context.JSON(http.StatusNotFound, model.Err{Error: "No card"})
		return
	}

	context.JSON(http.StatusOK, allName)
}
