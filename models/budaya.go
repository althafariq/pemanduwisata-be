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

func (b *BudayaModels) GetAllBudaya() ([]Budaya, error) {
	statement := `SELECT * FROM budaya`
	rows, err := b.db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	budayas := []Budaya{}
	for rows.Next() {
		var budaya Budaya
		err = rows.Scan(&budaya.ID, &budaya.Name, &budaya.DestinationID)
		if err != nil {
			return nil, err
		}
		budayas = append(budayas, budaya)
	}
	return budayas, nil
}

//function to get budaya by destination id	
func (b *BudayaModels) GetBudayaByDestinationID(id int) ([]Budaya, error) {
	statement := `SELECT * FROM budaya WHERE destination_id = ?`
	rows, err := b.db.Query(statement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	budayas := []Budaya{}
	for rows.Next() {
		var budaya Budaya
		err = rows.Scan(&budaya.ID, &budaya.Name, &budaya.DestinationID)
		if err != nil {
			return nil, err
		}
		budayas = append(budayas, budaya)
	}
	return budayas, nil
}