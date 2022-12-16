package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type BudayaModels struct {
	db *sql.DB
}

func NewBudayaModels(db *sql.DB) *BudayaModels {
	return &BudayaModels{
		db: db,
	}
}
