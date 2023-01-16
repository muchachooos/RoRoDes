package handler

import "github.com/jmoiron/sqlx"

type DataBase struct {
	DB *sqlx.DB
}
