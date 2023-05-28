package service

import "RoRoDes/model"

func (s *Service) InitGame(user string) (string, error) {
	return s.Storage.InitGameInDB(user)
}

func (s *Service) GetGame(id string) ([]model.FieldResponse, error) {
	return s.Storage.GetGameFromDB(id)
}

func (s *Service) GetGameId() ([]string, error) {
	return s.Storage.GetAllGameIdFromDB()
}
