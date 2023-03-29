package model

type Err struct {
	Error string `json:"error"`
}

type Config struct {
	DataSourceName string `json:"dataSourceName"`
	Port           int    `json:"port"`
	Key            string `json:"auth_key"`
}

type Card struct {
	CardID  string  `db:"card_id"`
	Name    *string `db:"name"`
	Damage  *int    `db:"damage"`
	Speed   *int    `db:"speed"`
	Health  *int    `db:"health"`
	Picture []byte  `db:"picture"`
}

type Unit struct {
	UnitID string  `json:"unit_id" db:"unit_id"`
	CardID string  `json:"card_id" db:"card_id"`
	GameId string  `json:"game_id" db:"game_id"`
	Name   *string `json:"name" db:"name"`
	Damage *int    `json:"damage" db:"damage"`
	Speed  *int    `json:"speed" db:"speed"`
	Health int     `json:"health" db:"health"`
}

type Field struct {
	Y      int     `json:"y" db:"y" db:"y"`
	X      int     `json:"x" db:"x" db:"x"`
	UnitID *string `json:"unit_id"   db:"unit_id"`
	GameID string  `json:"game_id" db:"game_id"`
}

type Body struct {
	Name   *string `json:"name" db:"name"`
	Damage *int    `json:"damage" db:"damage"`
	Speed  *int    `json:"speed" db:"speed"`
	Health *int    `json:"health" db:"health"`
}
