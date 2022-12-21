package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type TelpDaruratModels struct {
	db *sql.DB
}

func NewTelpDaruratModels(db *sql.DB) *TelpDaruratModels {
	return &TelpDaruratModels{
		db: db,
	}
}

func (t *TelpDaruratModels) GetAllTelpDarurat() ([]TelpDarurat, error) {
	statement := `SELECT * FROM telp_darurat`
	rows, err := t.db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	telpDarurats := []TelpDarurat{}
	for rows.Next() {
		var telpDarurat TelpDarurat
		err = rows.Scan(&telpDarurat.ID, &telpDarurat.Name, &telpDarurat.Number)
		if err != nil {
			return nil, err
		}
		telpDarurats = append(telpDarurats, telpDarurat)
	}
	return telpDarurats, nil
}
