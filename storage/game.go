package storage

import (
	"GameAPI/model"
	"github.com/google/uuid"
)

const countOfFields = 40

func (s *Storage) InitGameInDB() (string, error) {
	gameId := uuid.NewString()

	_, err := s.DB.Exec("INSERT INTO game (`game_id`) VALUES (?)", gameId)
	if err != nil {
		return "", err
	}

	for i := 1; i <= countOfFields; i++ {
		_, err = s.DB.Exec("INSERT INTO fields (`number`, `game_id`) VALUES (?,?)", i, gameId)
		if err != nil {
			return "", err
		}
	}

	return gameId, nil
}

func (s *Storage) GetGameFromDB(gameID string) ([]model.Field, error) {
	var game []model.Field

	err := s.DB.Select(&game, "SELECT `number`,`unit_id` FROM fields WHERE `game_id` = ?", gameID)
	if err != nil {
		return nil, err
	}

	return game, nil
}
