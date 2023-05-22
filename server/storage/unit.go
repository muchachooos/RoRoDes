package storage

import (
	"RoRoDes/model"
	"RoRoDes/utilities"
	"errors"
	"github.com/google/uuid"
)

var ErrMoveUnit = errors.New("unit cannot move in the specified direction")
var ErrCreateUnit = errors.New("unit cannot be created - this field is taken")

func (s *Storage) CreateUnitInDB(cardID, gameID string, x, y int) ([]model.UnitResponse, error) {
	var field []model.Field

	err := s.DB.Select(&field, "SELECT `unit_id` FROM field WHERE `x` = ? AND `y` = ? AND `game_id` = ?", x, y, gameID)
	if err != nil {
		return nil, err
	}

	if field[0].UnitID != nil {
		return nil, ErrCreateUnit
	}

	var card []model.Card

	err = s.DB.Select(&card, "SELECT `card_id`, `name`, `damage`, `speed`, `health`, `picture` FROM card WHERE `card_id` = ?", cardID)
	if err != nil {
		return nil, err
	}

	unitID := uuid.NewString()

	_, err = s.DB.Exec("INSERT INTO unit (`unit_id`, `card_id`,`game_id`, `name`, `damage`, `speed`, `health`) VALUES (?,?,?,?,?,?,?)", unitID, cardID, gameID, card[0].Name, card[0].Damage, card[0].Speed, card[0].Health)
	if err != nil {
		return nil, err
	}

	res, err := s.DB.Exec("UPDATE field SET unit_id = ? WHERE x = ? AND y = ? AND game_id = ?", unitID, x, y, gameID)
	if err != nil {
		return nil, err
	}

	countOfModifiedRows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if countOfModifiedRows == 0 {
		return nil, nil
	}

	var unit []model.UnitResponse

	err = s.DB.Select(&unit, "SELECT `unit_id`, `card_id`, `game_id`, `name`, `damage`, `speed`, `health` FROM unit WHERE `unit_id` = ?", unitID)
	if err != nil {
		return nil, err
	}

	unit[0].X = x
	unit[0].Y = y

	return unit, nil
}

func (s *Storage) GetUnitFromDB(unitID string) ([]model.UnitResponse, error) {
	var field []model.Field

	err := s.DB.Select(&field, "SELECT * FROM field WHERE `unit_id` = ?", unitID)
	if err != nil {
		return nil, err
	}

	var unitField []model.UnitResponse

	err = s.DB.Select(&unitField, "SELECT `unit_id`, `card_id`, `game_id`, `name`, `damage`, `speed`, `health` FROM unit WHERE `unit_id` = ?", unitID)
	if err != nil {
		return nil, err
	}

	unitField[0].X = field[0].X
	unitField[0].Y = field[0].Y

	return unitField, nil
}

func (s *Storage) MoveUnitInDB(unitID, dir string) (bool, error) {
	var unit model.Field

	err := s.DB.Get(&unit, "SELECT * FROM field WHERE `unit_id` = ?", unitID)
	if err != nil {
		return false, err
	}

	var fields []model.Field

	err = s.DB.Select(&fields, "SELECT * FROM field WHERE `game_id` = ?", unit.GameID)
	if err != nil {
		return false, err
	}

	gameMap := [][]*model.Field{
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
	}

	for i := range fields {
		y := fields[i].Y
		x := fields[i].X
		gameMap[y][x] = &fields[i]
	}

	newField, ok := utilities.CheckMove(unit.X, unit.Y, dir, gameMap)
	if ok == false {
		return false, ErrMoveUnit
	}

	if ok {
		result, err := s.DB.Exec("UPDATE field SET `unit_id` = ? WHERE `y` =? AND `x` = ?", nil, unit.Y, unit.X)
		if err != nil {
			return false, err
		}

		countOfModifiedRows, err := result.RowsAffected()
		if err != nil {
			return false, err
		}

		if countOfModifiedRows == 0 {
			return false, errors.New("failed set data")
		}

		res, err := s.DB.Exec("UPDATE field SET `unit_id` = ? WHERE `y` =? AND `x` = ?", unit.UnitID, newField.Y, newField.X)
		if err != nil {
			return false, err
		}

		countOfModifiedRows, err = res.RowsAffected()
		if err != nil {
			return false, err
		}

		if countOfModifiedRows == 0 {
			return false, errors.New("failed new field set data")
		}

	}

	return true, nil
}
