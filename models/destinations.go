package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DestinationModels struct {
	db *sql.DB
}

func NewDestinationModels(db *sql.DB) *DestinationModels {
	return &DestinationModels{
		db: db,
	}
}

func (d *DestinationModels) GetDestination() ([]Destination, error) {
	statement := "SELECT * FROM destination"
	rows, err := d.db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	destinations := []Destination{}
	for rows.Next() {
		var destination Destination
		err = rows.Scan(&destination.id, &destination.name, &destination.location)
		if err != nil {
			return nil, err
		}
		destinations = append(destinations, destination)
	}
	return destinations, nil
}