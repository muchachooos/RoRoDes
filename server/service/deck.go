package service

import "RoRoDes/model"

func (s *Service) AddCard(userLogin, cardId string) error {
	return s.Storage.AddCardInDB(userLogin, cardId)
}

func (s *Service) GetDeck(userLogin string) ([]model.CardResponse, error) {
	return s.Storage.GetDeckFromDB(userLogin)
}

func (s *Service) DeleteCard(userLogin, cardId string) (bool, error) {
	return s.Storage.DeleteCardInDB(userLogin, cardId)
}
