package handler

import (
	"RoRoDes/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) GetDeckHandler(context *gin.Context) {
	userLogin, ok := context.GetQuery("login")
	if userLogin == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Login is missing"})
		return
	}

	cardInDeck, err := s.Service.GetDeck(userLogin)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	if len(cardInDeck) == 0 {
		context.JSON(http.StatusNotFound, model.Err{Error: "No deck with this login: " + userLogin})
		return
	}

	context.JSON(http.StatusOK, cardInDeck)
}

func (s *Server) AddCardInDeckHandler(context *gin.Context) {
	userLogin, ok := context.GetQuery("login")
	if userLogin == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Login is missing"})
		return
	}

	cardId, ok := context.GetQuery("card_id")
	if cardId == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Login is missing"})
		return
	}

	ok, err := s.Service.AddCard(userLogin, cardId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, ok)
}
