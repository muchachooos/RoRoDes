package storage

import (
	"GameAPI/model"
	"GameAPI/utilities"
	"github.com/google/uuid"
)

const countOfFields = 40

func (s *Storage) InitGameInDB() (string, error) {
	gameId := uuid.NewString()

	_, err := s.DB.Exec("INSERT INTO game (`game_id`) VALUES (?)", gameId)
	if err != nil {
		return "", err
	}

	fieldsArray := utilities.CreateFields(gameId)

	for i := 0; i < countOfFields; i++ {
		_, err = s.DB.Exec("INSERT INTO field (`y`,`x`,`game_id`) VALUES (?,?,?)", fieldsArray[i].Y, fieldsArray[i].X, fieldsArray[i].GameID)
		if err != nil {
			return "", err
		}
	}

	return gameId, nil
}

func (s *Storage) GetGameFromDB(gameID string) ([]model.Field, error) {
	var game []model.Field

	err := s.DB.Select(&game, "SELECT `x`,`y`,`unit_id` FROM field WHERE `game_id` = ?", gameID)
	if err != nil {
		return nil, err
	}

	return game, nil
}

/*
const countOfFields = 40

		for i := 1; i <= countOfFields; i++ {
		_, err = s.DB.Exec("INSERT INTO field (`x`, `y`, `game_id`) VALUES (?,?,?)", i, gameId)
		if err != nil {
			return "", err
		}
	}
*/
