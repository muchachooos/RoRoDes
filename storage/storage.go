package storage

import "github.com/jmoiron/sqlx"

type Storage struct {
	DB *sqlx.DB
}
