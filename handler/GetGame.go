package handler

import (
	"GameAPI/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
)

//func (s *DataBase) GetGameHandler(context *gin.Context) {
//	gameID, ok := context.GetQuery("game_ID")
//	if gameID == "" || !ok {
//		context.Writer.WriteString("Game ID is missing")
//		return
//	}
//
//	result, err := s.DB.Exec("INSERT INTO game (`ID`,`GAME`)VALUES (?)", gameID)
//	if err != nil {
//		panic(err)
//	}
//}

func (s *DataBase) CreateGame(context *gin.Context) {
	bodyOnBytes, err := io.ReadAll(context.Request.Body)
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

	err = json.Unmarshal(bodyOnBytes, &game)
	if err != nil {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Unmarshal request body error: " + err.Error()})
	}

	id := uuid.NewString()

	_, err = s.DB.Exec("INSERT INTO game (`id`) VALUES (?)", id)
	if err != nil {
		panic(err)
	}

	for i := range game {
		query := fmt.Sprintf("UPDATE game SET `field_%d` = ? WHERE `id` = ?", game[i].Number)
		_, err = s.DB.Exec(query, game[i].Unit, id)
		if err != nil {
			panic(err)
		}
	}
}
