package service

import "RoRoDes/model"

func (s *Service) GetCard(cardID string) ([]model.CardResponse, error) {
	return s.Storage.GetCardFromDB(cardID)
}
