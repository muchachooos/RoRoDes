package storage

import (
	"RoRoDes/model"
	"RoRoDes/utilities"
	"github.com/google/uuid"
)

const countOfFields = 40

func (s *Storage) InitGameInDB(user string) (string, error) {
	gameId := uuid.NewString()

	_, err := s.DB.Exec("INSERT INTO game (`game_id`, `first_user`) VALUES (?, ?)", gameId, user)
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

func (s *Storage) GetGameFromDB(gameID string) ([]model.FieldResponse, error) {
	var game []model.FieldResponse

	err := s.DB.Select(&game, "SELECT `x`,`y`,`unit_id` FROM field WHERE `game_id` = ?", gameID)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (s *Storage) GetAllGameIdFromDB() ([]string, error) {
	var allId []string

	err := s.DB.Select(&allId, "SELECT `game_id` FROM game")
	if err != nil {
		return nil, err
	}

	return allId, nil
}
