package service

import "RoRoDes/model"

func (s *Service) InitGame() (string, error) {
	return s.Storage.InitGameInDB()
}

func (s *Service) GetGame(id string) ([]model.FieldResponse, error) {
	return s.Storage.GetGameFromDB(id)
}
