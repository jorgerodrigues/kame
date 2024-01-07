package models

import "database/sql"

type Model struct {
	User *UserModel
}

func New(db *sql.DB) *Model {
	return &Model{
		User: &UserModel{
			DB: db,
		},
	}
}
