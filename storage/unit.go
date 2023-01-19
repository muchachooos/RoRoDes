package storage

import (
	"GameAPI/model"
	"github.com/google/uuid"
)

func (s *Storage) CreateUnitInDB(cardID string) ([]model.Unit, error) {
	var card []model.Card

	err := s.DB.Select(&card, "SELECT `id`, `name`, `damage`, `speed`, `health` FROM card WHERE `id` = ?", cardID)
	if err != nil {
		return nil, err
	}

	unitID := uuid.NewString()

	_, err = s.DB.Exec("INSERT INTO unit (`id`, `card_id`,`name`,`damage`,`speed`,`health`) VALUES (?,?,?,?,?,?)", unitID, cardID, card[0].Name, card[0].Damage, card[0].Speed, card[0].Health)
	if err != nil {
		return nil, err
	}

	var unit []model.Unit

	err = s.DB.Select(&unit, "SELECT `id`, `card_id`, `name`, `damage`, `speed`, `health` FROM unit WHERE `id` = ?", unitID)
	if err != nil {
		return nil, err
	}
	return unit, nil
}
