package service

import "RoRoDes/model"

func (s *Service) CreateUnit(cardID, gameID string, x, y int) ([]model.UnitResponse, error) {
	return s.Storage.CreateUnitInDB(cardID, gameID, x, y)
}

func (s *Service) GetUnit(id string) ([]model.UnitResponse, error) {
	return s.GetUnit(id)
}

func (s *Service) MoveUnit(id, dir string) (bool, error) {
	return s.Storage.MoveUnitInDB(id, dir)
}
