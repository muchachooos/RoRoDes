package utilities

import (
	"RoRoDes/model"
)

const (
	Up    = "UP"
	Down  = "DOWN"
	Left  = "LEFT"
	Right = "RIGHT"
)

func CheckMove(x, y int, direction string, game [][]*model.Field) (newField *model.Field, ok bool) {
	switch direction {
	case Up:
		field := game[y-1][x]
		if field.UnitID != nil {
			return nil, false
		}
		return field, true
	case Down:
		field := game[y+1][x]
		if field.UnitID != nil {
			return nil, false
		}
		return field, true
	case Left:
		field := game[y][x-1]
		if field.UnitID != nil {
			return nil, false
		}
		return field, true
	case Right:
		field := game[y][x+1]
		if field.UnitID != nil {
			return nil, false
		}
		return field, true
	default:
		return nil, false
	}
}
