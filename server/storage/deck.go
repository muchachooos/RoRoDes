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

func (s *Storage) DeleteCardFromDB(userLogin, cardID string) error {
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

	// Проверяем на наличие такой карты в колоде пользователя
	err = s.DB.Get(&cardInDeck, "SELECT * FROM card_in_deck WHERE `deck_id` = ? AND `card_id` = ?", deckID, cardID)
	if err != nil {
		// Если карты с таким ID нет, возвращаем ошибку
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrCardNotFound
		}
		return err
	}

	// Если такая карта одна, то удаляем ее
	if cardInDeck.Count == 1 {
		_, err = s.DB.Exec("DELETE FROM card_in_deck WHERE `deck_id` = ? AND `card_id` = ?", deckID, cardID)
		return err
	}

	// Если такая карта не одна, то уменьшаем её кол-во на 1
	_, err = s.DB.Exec("UPDATE card_in_deck SET `count` = `count` - 1 WHERE `deck_id` = ? AND `card_id` = ?", deckID, cardID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetDeckFromDB(userLogin string) ([]model.CardResponse, error) {
	var deckID string

	// Получаем ID колоды пользователя
	err := s.DB.Get(&deckID, "SELECT `id` FROM deck WHERE `user_login` = ?", userLogin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrDeckNotFound
		}
		return nil, err
	}

	var cardsInDeck []model.CardInDeck

	// Получаем все карты в колоде пользователя
	err = s.DB.Select(&cardsInDeck, "SELECT * FROM card_in_deck WHERE `deck_id` = ?", deckID)
	if err != nil {
		return nil, err
	}

	//Создаём слайс и заполняем его ID карт из колоды
	cardIDs := make([]string, 0, len(cardsInDeck))

	for i := range cardsInDeck {
		cardIDs = append(cardIDs, cardsInDeck[i].CardID)
	}

	query, args, err := sqlx.In("SELECT `card_id`, `name`, `name`, `speed`, `health` FROM card WHERE `card_id` IN (?)", cardIDs)
	if err != nil {
		return nil, err
	}

	var cards []model.CardResponse

	//Находим все карты с нужными ID
	err = s.DB.Select(&cards, query, args...)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

/*
var cardIDs []model.Card

err = s.DB.Select(&cardIDs, "SELECT * FROM card WHERE `card_id` IN (SELECT card_id FROM card_in_deck WHERE `deck_id` = ?)", deckID)
if err != nil {
return nil, err
}
*/

/*
func (s *Storage) GetDeckFromDB(userLogin string) ([]model.CardResponse, error) {
	var deckID string

	// Получаем ID колоды пользователя
	err := s.DB.Get(&deckID, "SELECT `id` FROM deck WHERE `user_login` = ?", userLogin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrDeckNotFound
		}
		return nil, err
	}

	var cardIDs []string

	err = s.DB.Select(&cardIDs, "SELECT * FROM card_in_deck WHERE `deck_id` = ?", deckID)
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
*/
