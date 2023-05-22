package utilities

import (
	"RoRoDes/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFields(t *testing.T) {
	id := "6e7a3251-3333-5555-7777-0a3739b1be0e"

	result := CreateFields(id)

	expected := []model.Field{
		{
			Y:      0,
			X:      0,
			GameID: id,
		},
		{
			Y:      0,
			X:      1,
			GameID: id,
		},
		{
			Y:      0,
			X:      2,
			GameID: id,
		},
		{
			Y:      0,
			X:      3,
			GameID: id,
		},
		{
			Y:      0,
			X:      4,
			GameID: id,
		},
		{
			Y:      0,
			X:      5,
			GameID: id,
		},
		{
			Y:      0,
			X:      6,
			GameID: id,
		},
		{
			Y:      0,
			X:      7,
			GameID: id,
		},
		{
			Y:      1,
			X:      0,
			GameID: id,
		},
		{
			Y:      1,
			X:      1,
			GameID: id,
		},
		{
			Y:      1,
			X:      2,
			GameID: id,
		},
		{
			Y:      1,
			X:      3,
			GameID: id,
		},
		{
			Y:      1,
			X:      4,
			GameID: id,
		},
		{
			Y:      1,
			X:      5,
			GameID: id,
		},
		{
			Y:      1,
			X:      6,
			GameID: id,
		},
		{
			Y:      1,
			X:      7,
			GameID: id,
		},
		{
			Y:      2,
			X:      0,
			GameID: id,
		},
		{
			Y:      2,
			X:      1,
			GameID: id,
		},
		{
			Y:      2,
			X:      2,
			GameID: id,
		},
		{
			Y:      2,
			X:      3,
			GameID: id,
		},
		{
			Y:      2,
			X:      4,
			GameID: id,
		},
		{
			Y:      2,
			X:      5,
			GameID: id,
		},
		{
			Y:      2,
			X:      6,
			GameID: id,
		},
		{
			Y:      2,
			X:      7,
			GameID: id,
		},
		{
			Y:      3,
			X:      0,
			GameID: id,
		},
		{
			Y:      3,
			X:      1,
			GameID: id,
		},
		{
			Y:      3,
			X:      2,
			GameID: id,
		},
		{
			Y:      3,
			X:      3,
			GameID: id,
		},
		{
			Y:      3,
			X:      4,
			GameID: id,
		},
		{
			Y:      3,
			X:      5,
			GameID: id,
		},
		{
			Y:      3,
			X:      6,
			GameID: id,
		},
		{
			Y:      3,
			X:      7,
			GameID: id,
		},
		{
			Y:      4,
			X:      0,
			GameID: id,
		},
		{
			Y:      4,
			X:      1,
			GameID: id,
		},
		{
			Y:      4,
			X:      2,
			GameID: id,
		},
		{
			Y:      4,
			X:      3,
			GameID: id,
		},
		{
			Y:      4,
			X:      4,
			GameID: id,
		},
		{
			Y:      4,
			X:      5,
			GameID: id,
		},
		{
			Y:      4,
			X:      6,
			GameID: id,
		},
		{
			Y:      4,
			X:      7,
			GameID: id,
		},
	}

	assert.ElementsMatch(t, expected, result)
}
