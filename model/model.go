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
	Health  int     `db:"health"`
	Picture []byte  `db:"picture"`
}

type Unit struct {
	UnitID   string  `json:"unit_id" db:"unit_id"`
	CardID   string  `json:"card_id" db:"card_id"`
	GameID   string  `json:"game_id" db:"game_id"`
	FieldNum int     `json:"field_num" db:"field_num"`
	Name     *string `json:"name" db:"name"`
	Damage   *int    `json:"damage" db:"damage"`
	Speed    *int    `json:"speed" db:"speed"`
	Health   int     `json:"health" db:"health"`
}

type Field struct {
	Number int     `json:"field_num" db:"number"`
	UnitID *string `json:"unit"   db:"unit_id"`
}

//type Game struct {
//	Fields []GameField `json:"game"`
//}

//type T struct {
//	Game []struct {
//		Field string  `json:"field"`
//		Unit  *string `json:"unit"`
//	} `json:"game"`
//}
