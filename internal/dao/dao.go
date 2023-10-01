package dao

import "github.com/jmoiron/sqlx"

type Dao struct {
	db *sqlx.DB
}

// NewDao returns a new Dao
func NewDao(db *sqlx.DB) *Dao {
	return &Dao{db}
}
