package models

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Model struct {
	User *UserModel
	URLs *URLModel
}

func New(db *pgxpool.Pool) *Model {
	return &Model{
		User: &UserModel{
			DB: db,
		},
		URLs: &URLModel{
			DB: db,
		},
	}
}
