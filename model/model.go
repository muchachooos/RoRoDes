package model

type Config struct {
	DataSourceName string `json:"dataSourceName"`
	Port           int    `json:"port"`
	Key            string `json:"auth_key"`
}

type GameField struct {
	Field int     `json:"field"`
	Unit  *string `json:"unit"`
}

type Game struct {
	Fields []GameField `json:"game"`
}

//type T struct {
//	Game []struct {
//		Field string  `json:"field"`
//		Unit  *string `json:"unit"`
//	} `json:"game"`
//}
