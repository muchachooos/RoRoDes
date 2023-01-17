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

	var fields []model.Field

	err = json.Unmarshal(bodyInBytes, &fields)
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

	context.JSON(http.StatusOK, gameId)
}

func (s *DataBase) CreateGameHandler(context *gin.Context) {
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

	var fields []model.Field

	err = json.Unmarshal(bodyOnBytes, &fields)
	if err != nil {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Unmarshal request body error: " + err.Error()})
	}

	gameId := uuid.NewString()

	query := `INSERT game VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
        ?, ?, ?);`

	var units []any
	for i := range fields {
		units = append(units, fields[i].Unit)
	}

	idInSlice := []any{gameId}

	var args = append(idInSlice, units...)

	_, err = s.DB.Exec(query, args...)
	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, gameId)
}

func (s *DataBase) GetGameHandler(context *gin.Context) {
	gameID, ok := context.GetQuery("id")
	if gameID == "" || !ok {
		context.Writer.WriteString("Game ID is missing")
		return
	}

	var game []model.Game

	err := s.DB.Select(&game, "SELECT * FROM game WHERE `id` = ?", gameID)
	if err != nil {
		panic(err)
	}

	if len(game) == 0 {
		context.Status(404)
		context.Writer.WriteString("No students with this ID")
		return
	}

	context.JSON(http.StatusOK, game)
}

/*

	id := uuid.NewString()
	idInSlice := []any{id}

	query := `INSERT game VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
        ?, ?, ?);`

	var units []any
	for i := range fields {
		units = append(units, fields[i].Unit)
	}

	var args = append(idInSlice, units...)

	_, err = s.DB.Exec(query, args...)
	if err != nil {
		panic(err)
	}




_, err = s.DB.Exec("INSERT INTO game (`id`) VALUES (?)", id)
	if err != nil {
		panic(err)
	}
for i := range game {
		if game[i].Unit != nil {
			query := fmt.Sprintf("UPDATE game SET `field_%d` = ? WHERE `id` = ?", game[i].Number)
			_, err = s.DB.Exec(query, game[i].Unit, id)
			if err != nil {
				panic(err)
			}
		}
	}
*/
