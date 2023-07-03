package storage

import (
	"RoRoDes/model"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

func (s *Storage) AddCardInDB(userLogin, cardID string) error {

	var deckID string

	// Получаем ID колоды пользователя
	err := s.DB.Get(&deckID, "SELECT `id` FROM deck WHERE `user_login` = ?", userLogin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrDeckNotFound
		}
		return err
	}

	var cardInDeck model.CardInDeck

	// Проверяем на наличие такой же карты в колоде пользователя
	err = s.DB.Get(&cardInDeck, "SELECT * FROM card_in_deck WHERE `deck_id` = ? AND `card_id` = ?", deckID, cardID)
	if err != nil {
		// Если карты с таким ID нет, то добавляем её в колличестве одной штуки
		if errors.Is(err, sql.ErrNoRows) {
			_, err = s.DB.Exec("INSERT INTO card_in_deck (`deck_id`, `card_id`, `count`) VALUES (?, ?, ?)", deckID, cardID, 1)
			return err
		}
		return err
	}

	// Если такая карта есть, то увеличиваем её кол-во на 1
	_, err = s.DB.Exec("UPDATE card_in_deck SET `count` = `count` + 1 WHERE `deck_id` = ? AND `card_id` = ?", deckID, cardID)
	if err != nil {
		return err
	}

	return nil
}

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

	query, args, err := sqlx.In("SELECT `card_id`, `name`, `damage`, `speed`, `health` FROM card WHERE `card_id` IN (?)", cardIDs)
	if err != nil {
		return nil, err
	}

	var cardInDeck []model.CardResponse

	err = s.DB.Select(&cardInDeck, query, args...)
	if err != nil {
		return nil, err
	}

	return cardInDeck, nil
}

func (s *Storage) DeleteCardInDB(userLogin, cardID string) (bool, error) {
	var deckID string

	err := s.DB.Get(&deckID, "SELECT `id` FROM deck WHERE `user_login` = ?", userLogin)
	if err != nil {
		return false, err
	}

	res, err := s.DB.Exec("DELETE FROM card_in_deck WHERE `deck_id` = ? AND `card_id` = ? LIMIT 1", deckID, cardID)
	if err != nil {
		return false, err
	}

	countOfDeletedRows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if countOfDeletedRows == 0 {
		return false, nil
	}

	return true, nil
}
