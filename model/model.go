package model

type Err struct {
	Error string `json:"error"`
}

type Config struct {
	DataSourceName string `json:"dataSourceName"`
	Port           int    `json:"port"`
	Key            string `json:"auth_key"`
}

type Field struct {
	Number int     `json:"number"`
	Unit   *string `json:"unit"`
}

type Unit struct {
	ID     string  `db:"id"`
	CardID *string `db:"card_id"`
	Health int     `db:"health"`
}

type Game struct {
	ID      string  `db:"id"`
	Field1  *string `db:"field_1"`
	Field2  *string `db:"field_2"`
	Field3  *string `db:"field_3"`
	Field4  *string `db:"field_4"`
	Field5  *string `db:"field_5"`
	Field6  *string `db:"field_6"`
	Field7  *string `db:"field_7"`
	Field8  *string `db:"field_8"`
	Field9  *string `db:"field_9"`
	Field10 *string `db:"field_10"`
	Field11 *string `db:"field_11"`
	Field12 *string `db:"field_12"`
	Field13 *string `db:"field_13"`
	Field14 *string `db:"field_14"`
	Field15 *string `db:"field_15"`
	Field16 *string `db:"field_16"`
	Field17 *string `db:"field_17"`
	Field18 *string `db:"field_18"`
	Field19 *string `db:"field_19"`
	Field20 *string `db:"field_20"`
	Field21 *string `db:"field_21"`
	Field22 *string `db:"field_22"`
	Field23 *string `db:"field_23"`
	Field24 *string `db:"field_24"`
	Field25 *string `db:"field_25"`
	Field26 *string `db:"field_26"`
	Field27 *string `db:"field_27"`
	Field28 *string `db:"field_28"`
	Field29 *string `db:"field_29"`
	Field30 *string `db:"field_30"`
	Field31 *string `db:"field_31"`
	Field32 *string `db:"field_32"`
	Field33 *string `db:"field_33"`
	Field34 *string `db:"field_34"`
	Field35 *string `db:"field_35"`
	Field36 *string `db:"field_36"`
	Field37 *string `db:"field_37"`
	Field38 *string `db:"field_38"`
	Field39 *string `db:"field_39"`
	Field40 *string `db:"field_40"`
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
