package storage

import (
	"RoRoDes/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func (s *Storage) GetDeckFromDB(userLogin string) ([]model.CardResponse, error) {
	var deckID string

	err := s.DB.Get(&deckID, "SELECT `id` FROM deck WHERE `user_login` = ?", userLogin)
	if err != nil {
		return nil, err
	}

	var cardIDs []string

	err = s.DB.Select(&cardIDs, "SELECT `card_id` FROM card_in_deck WHERE `deck_id` = ?", deckID)
	if err != nil {
		return nil, err
	}

	var cardInDeck []model.CardResponse

	query, args, err := sqlx.In("SELECT `card_id`, `name`, `damage`, `speed`, `health` FROM card WHERE `card_id` IN (?)", cardIDs)
	if err != nil {
		return nil, err
	}

	fmt.Println(query)
	fmt.Println(args)

	err = s.DB.Select(&cardInDeck, query, args...)
	if err != nil {
		return nil, err
	}

	return cardInDeck, nil
}

func (s *Storage) AddCardInDB(userLogin, cardId string) (bool, error) {
	var deckID string

	err := s.DB.Get(&deckID, "SELECT `id` FROM deck WHERE `user_login` = ?", userLogin)
	if err != nil {
		return false, err
	}

	_, err = s.DB.Exec("INSERT INTO card_in_deck (`deck_id`, `card_id`) VALUES (?, ?)", deckID, cardId)
	if err != nil {
		return false, err
	}

	return true, nil
}
