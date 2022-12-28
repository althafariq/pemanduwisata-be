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

func (d *DestinationModels) GetAllDestinations() ([]Destination, error) {
	statement := `SELECT * FROM destinations`
	rows, err := d.db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	destinations := []Destination{}
	for rows.Next() {
		var destination Destination
		err = rows.Scan(
			&destination.ID, 
			&destination.Name, 
			&destination.Location, 
			&destination.Description,
			&destination.BudayaName,
			&destination.BudayaDescription,
			&destination.PhotoPath,
			)
		if err != nil {
			return nil, err
		}
		destinations = append(destinations, destination)
	}
	return destinations, nil
}

func (d *DestinationModels) GetDestinationbyID(id int) ([]Destination, error) {
	statement := `SELECT * FROM destinations WHERE id = ?`
	rows, err := d.db.Query(statement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	destinations := []Destination{}
	for rows.Next() {
		var destination Destination
		err = rows.Scan(
			&destination.ID,
			&destination.Name,
			&destination.Location,
			&destination.Description,
			&destination.BudayaName,
			&destination.BudayaDescription,
			&destination.PhotoPath,
			)
		if err != nil {
			return nil, err
		}
		destinations = append(destinations, destination)
	}
	return destinations, nil
}

func (d *DestinationModels) CreateDestination(destination Destination) (int, error) {
	statement := `INSERT INTO destinations (name, location, description, budaya_name, budaya_description, photo_path) VALUES (?, ?, ?, ?, ?, "media/destination/default.jpg")`
	result, err := d.db.Exec(statement, 
		destination.Name, 
		destination.Location, 
		destination.Description,
		destination.BudayaName,
		destination.BudayaDescription,
		destination.PhotoPath,
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

//function to upload image to destinations table
func (d *DestinationModels) InsertDestinationImage(ID int, PhotoPath string) error {
	statement := `UPDATE destinations SET photo_path = ? WHERE id = ?`

	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, error := tx.Exec(statement, PhotoPath, ID)
	if error != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (d *DestinationModels) UpdateDestination(ID int, Name, Location, Description, BudayaName, BudayaDescription, PhotoPath string) error {
	statement := `UPDATE destinations SET name = ?, location = ?, description = ?, budaya_name = ?, budaya_description = ?, photo_path = ? WHERE id = ?`
	_, err := d.db.Exec(statement,
		Name,
		Location,
		Description,
		BudayaName,
		BudayaDescription,
		PhotoPath,
		ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *DestinationModels) DeleteDestination(id int) error {
	statement := `DELETE FROM destinations WHERE id = ?`
	_, err := d.db.Exec(statement, id)
	if err != nil {
		return err
	}
	return nil
}
