package storage

import "RoRoDes/model"

func (s *Storage) GetCardFromDB(cardID string) ([]model.CardResponse, error) {
	var characteristics []model.CardResponse

	err := s.DB.Select(&characteristics, "SELECT `card_id`, `name`, `damage`, `speed`, `health` FROM card WHERE `card_id` = ?", cardID)
	if err != nil {
		return nil, err
	}

	return characteristics, nil
}

func (s *Storage) GetAllNameFromDB() ([]string, error) {
	var allName []string

	err := s.DB.Select(&allName, "SELECT `name` FROM card")
	if err != nil {
		return nil, err
	}

	return allName, nil
}
