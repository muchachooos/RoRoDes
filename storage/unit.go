package storage

import (
	"GameAPI/model"
	"GameAPI/utilities"
	"errors"
	"github.com/google/uuid"
)

var ErrMoveUnit = errors.New("unit cannot move in the specified direction")

func (s *Storage) CreateUnitInDB(cardID, gameID string, x, y int) ([]model.Unit, error) {
	var card []model.Card

	err := s.DB.Select(&card, "SELECT `card_id`, `name`, `damage`, `speed`, `health`, `picture` FROM card WHERE `card_id` = ?", cardID)
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

	var unit []model.Unit

	err = s.DB.Select(&unit, "SELECT `unit_id`, `card_id`, `game_id`, `name`, `damage`, `speed`, `health` FROM unit WHERE `unit_id` = ?", unitID)
	if err != nil {
		return nil, err
	}

	return unit, nil
}

func (s *Storage) GetUnitFromDB(unitID string) ([]model.Field, error) {
	var unitField []model.Field

	err := s.DB.Select(&unitField, "SELECT * FROM field WHERE `unit_id` = ?", unitID)
	if err != nil {
		return nil, err
	}

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

//var gameId string
//
//err := s.DB.Get(&gameId, "SELECT `game_id` FROM unit WHERE `unit_id` = ?", unitID)
//if err != nil {
//	return false
//}

/*
	res, err := s.DB.Exec("UPDATE field SET `y` = ?, `x` = ?, `unit_id` = ?, `game_id` = ? WHERE `y` =? AND `x` = ?", newField.Y, newField.X, unit.UnitID, unit.GameID, unit.Y, unit.X)

	res, err := s.DB.Exec("UPDATE field SET `x` = newField.X, `y` = newField.Y WHERE `unit_id` = ? AND `game_id` = ?", unit[0].UnitID, unit[0].GameID)
	if err != nil {
		return nil, err
	}

*/
//gameMap[unit.Y][unit.X+1]
//case Right:
//field := game[y][x+1]
//if field.UnitID != nil {
//return false
//}

/*
	var unit []model.Field

	err := s.DB.Select(&unit, "SELECT * FROM field WHERE `unit_id` = ?", unitId)
	if err != nil {
		return nil, err
	}

	var fields []*model.Field

	err = s.DB.Select(&fields, "SELECT * FROM field WHERE `game_id` = ?", unit[0].GameID)
	if err != nil {
		return nil, err
	}

	field := utilities.TransformArray(fields)

	return nil, err
*/

/*
	fmt.Println(game[0][0], game[0][1], game[0][2], game[0][3], game[0][4], game[0][5], game[0][6], game[0][7])
	fmt.Println("----------------------")
	fmt.Println(game[1][0], game[1][1], game[1][2], game[1][3], game[1][4], game[1][5], game[1][6], game[1][7])
	fmt.Println("----------------------------")
	fmt.Println(game[2][0], game[2][1], game[2][2], game[2][3], game[2][4], game[2][5], game[2][6], game[2][7])
	fmt.Println("----------------------------------")
	fmt.Println(game[3][0], game[3][1], game[3][2], game[3][3], game[3][4], game[3][5], game[3][6], game[3][7])
	fmt.Println("----------------------------------------")
	fmt.Println(game[4][0], game[4][1], game[4][2], game[4][3], game[4][4], game[4][5], game[4][6], game[4][7])
*/

/*
	err = s.DB.Select(&game[0], "SELECT * FROM field WHERE `game_id` = ? AND `y` = ?", unit[0].GameID, 0)
	if err != nil {
		return nil, err
	}

	err = s.DB.Select(&game[1], "SELECT * FROM field WHERE `game_id` = ? AND `y` = ?", unit[0].GameID, 1)
	if err != nil {
		return nil, err
	}

	err = s.DB.Select(&game[2], "SELECT * FROM field WHERE `game_id` = ? AND `y` = ?", unit[0].GameID, 2)
	if err != nil {
		return nil, err
	}

	err = s.DB.Select(&game[3], "SELECT * FROM field WHERE `game_id` = ? AND `y` = ?", unit[0].GameID, 3)
	if err != nil {
		return nil, err
	}

	err = s.DB.Select(&game[4], "SELECT * FROM field WHERE `game_id` = ? AND `y` = ?", unit[0].GameID, 4)
	if err != nil {
		return nil, err
	}

	newField, move := utilities.CheckMove(unit[0].X, unit[0].Y, dir, game)
	if move == false {
		return nil, err
	}

	if newField == nil {
		return nil, err
	}

	res, err := s.DB.Exec("UPDATE field SET `x` = newField.X, `y` = newField.Y WHERE `unit_id` = ? AND `game_id` = ?", unit[0].UnitID, unit[0].GameID)
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

	var unitField []model.Field

	err = s.DB.Select(&unitField, "SELECT * FROM field WHERE `unit_id` = ?", unit[0].UnitID)
	if err != nil {
		return nil, err
	}

	return unitField, nil
*/

/*
	var unitField [40]*model.Field

	err := s.DB.Select(&unitField, "SELECT * FROM field WHERE `unit_id` = ?", unitId)
	if err != nil {
		return nil, err
	}

	game := utilities.TransformArray(unitField)

	res := utilities.CheckMove(unitField[0].X, unitField[0].Y, dir, game)
	if res == false {
		return nil, err
	}

	return nil, err
*/
