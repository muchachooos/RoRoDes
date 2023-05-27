package service

import "RoRoDes/model"

func (s *Service) GetDeck(userLogin string) ([]model.CardResponse, error) {
	return s.Storage.GetDeckFromDB(userLogin)
}
