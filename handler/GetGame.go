package handler

import (
	"GameAPI/model"
	"encoding/json"
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

	query := `INSERT INTO game (id, field_1, field_2, field_3, field_4, field_5, field_6, field_7, field_8, field_9, field_10,
                  field_11, field_12, field_13, field_14, field_15, field_16, field_17, field_18, field_19, field_20,
                  field_21, field_22, field_23, field_24, field_25, field_26, field_27, field_28, field_29, field_30,
                  field_31, field_32, field_33, field_34, field_35, field_36, field_37, field_38, field_39, field_40)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
        ?, ?, ?);`

	_, err = s.DB.Exec(query, id)
	if err != nil {
		panic(err)
	}

}

/*










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
