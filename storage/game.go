package storage

import (
	"github.com/google/uuid"
)

const countOfFields = 40

func (s *Storage) InitGameInDB() (string, error) {
	gameId := uuid.NewString()

	_, err := s.DB.Exec("INSERT INTO game (`id`) VALUES (?)", gameId)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= countOfFields; i++ {
		_, err = s.DB.Exec("INSERT INTO fields (`number`, `game_id`) VALUES (?,?)", i, gameId)
		if err != nil {
			panic(err)
		}
	}

	return gameId, err
}
