package model

type Err struct {
	Error string `json:"error"`
}

type Config struct {
	DataSourceName string `json:"dataSourceName"`
	Port           int    `json:"port"`
	Key            string `json:"auth_key"`
}

type Unit struct {
	ID     string  `db:"id"`
	CardID *string `db:"card_id"`
	Health int     `db:"health"`
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
