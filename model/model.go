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
	ID      string  `db:"id"`
	Name    *string `db:"name"`
	Damage  *int    `db:"damage"`
	Speed   *int    `db:"speed"`
	Health  int     `db:"health"`
	Picture byte    `db:"picture"`
}

type Unit struct {
	ID     string  `json:"id" db:"id"`
	CardID string  `json:"card_id" db:"card_id"`
	Name   *string `json:"name" db:"name"`
	Damage *int    `json:"damage" db:"damage"`
	Speed  *int    `json:"speed" db:"speed"`
	Health int     `json:"health" db:"health"`
}

type Field struct {
	Number int     `json:"number" db:"number"`
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
