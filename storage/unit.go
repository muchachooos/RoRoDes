package storage

import (
	"GameAPI/model"
	"github.com/google/uuid"
)

func (s *Storage) CreateUnitInDB(cardID, gameID string, fieldNum int) ([]model.Unit, error) {
	var card []model.Card

	err := s.DB.Select(&card, "SELECT `card_id`, `name`, `damage`, `speed`, `health`, `picture` FROM card WHERE `card_id` = ?", cardID)
	if err != nil {
		return nil, err
	}

	unitID := uuid.NewString()

	_, err = s.DB.Exec("INSERT INTO unit (`unit_id`, `card_id`,`game_id`, `field_num`, `name`, `damage`, `speed`, `health`) VALUES (?,?,?,?,?,?,?,?)", unitID, cardID, gameID, fieldNum, card[0].Name, card[0].Damage, card[0].Speed, card[0].Health)
	if err != nil {
		return nil, err
	}

	var unit []model.Unit

	err = s.DB.Select(&unit, "SELECT `unit_id`, `card_id`, `game_id`, `field_num`, `name`, `damage`, `speed`, `health` FROM unit WHERE `unit_id` = ?", unitID)
	if err != nil {
		return nil, err
	}
	return unit, nil
}

func (s *Storage) GetUnitFromDB(unitID string) ([]model.Unit, error) {
	var unit []model.Unit

	err := s.DB.Select(&unit, "SELECT * FROM unit WHERE `unit_id` = ?", unitID)
	if err != nil {
		return nil, err
	}

	return unit, err
}
