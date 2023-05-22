package utilities

import (
	"RoRoDes/model"
)

func CreateFields(gameId string) [40]model.Field {
	var fields [40]model.Field

	x := 0
	for i := range fields {
		fields[i].GameID = gameId

		fields[i].X = x
		x++
		if x == 8 {
			x = 0
		}
	}

	y := 0
	for i := range fields {
		fields[i].Y = y
		y++
		if y == 5 {
			y = 0
		}

	}

	return fields
}
