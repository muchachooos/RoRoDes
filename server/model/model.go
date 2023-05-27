package model

type Err struct {
	Error string `json:"error"`
}

type Config struct {
	Port   int    `json:"port"`
	DBConf DBConf `json:"DataBase"`
}

type DBConf struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	DBName   string `json:"dataBaseName"`
	DBPort   int    `json:"db_port"`
}

type Card struct {
	CardID  string  `json:"card_id" db:"card_id"`
	Name    *string `json:"name"   db:"name"`
	Damage  *int    `json:"damage" db:"damage"`
	Speed   *int    `json:"speed"  db:"speed"`
	Health  *int    `json:"health" db:"health"`
	Picture []byte  `json:"picture" db:"picture"`
}

type CardResponse struct {
	CardID string  `json:"card_id" db:"card_id"`
	Name   *string `json:"name"   db:"name"`
	Damage *int    `json:"damage" db:"damage"`
	Speed  *int    `json:"speed"  db:"speed"`
	Health *int    `json:"health" db:"health"`
}

type Unit struct {
	UnitID string  `json:"unit_id" db:"unit_id"`
	CardID string  `json:"card_id" db:"card_id"`
	GameId string  `json:"game_id" db:"game_id"`
	Name   *string `json:"name"  db:"name"`
	Damage *int    `json:"damage" db:"damage"`
	Speed  *int    `json:"speed" db:"speed"`
	Health int     `json:"health" db:"health"`
}

type UnitResponse struct {
	Y      int     `json:"y" db:"y" db:"y"`
	X      int     `json:"x" db:"x" db:"x"`
	UnitID *string `json:"unit_id"   db:"unit_id"`
	CardID string  `json:"card_id" db:"card_id"`
	GameId string  `json:"game_id" db:"game_id"`
	Name   *string `json:"name"  db:"name"`
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

type FieldResponse struct {
	Y      int     `json:"y" db:"y" db:"y"`
	X      int     `json:"x" db:"x" db:"x"`
	UnitID *string `json:"unit_id"   db:"unit_id"`
}

type Body struct {
	Name   *string `json:"name" db:"name"`
	Damage *int    `json:"damage" db:"damage"`
	Speed  *int    `json:"speed" db:"speed"`
	Health *int    `json:"health" db:"health"`
}
