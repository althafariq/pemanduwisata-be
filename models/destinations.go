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

//function to get all destinations including review, budaya, and photos entities
func (d *DestinationModels) GetAllDestinations() ([]Destination, error) {
	statement := `SELECT 
	d.*, b.name, p.path
	 FROM destinations d
	 JOIN budaya b ON d.id = b.destination_id
	 JOIN photos p ON d.id = p.destination_id`
	rows, err := d.db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	destinations := []Destination{}
	for rows.Next() {
		var destination Destination
		err = rows.Scan(&destination.ID, &destination.Name, &destination.Location, &destination.Description, &destination.BudayaName, &destination.Photo_path)
		if err != nil {
			return nil, err
		}
		destinations = append(destinations, destination)
	}
	return destinations, nil
}

//function to get destination by id including review, budaya, and photos entities
func (d *DestinationModels) GetDestinationbyID(id int) ([]Destination, error) {
	statement := `SELECT 
	d.*, b.name, p.path
	 FROM destinations d
	 JOIN budaya b ON d.id = b.destination_id
	 JOIN photos p ON d.id = p.destination_id
	 WHERE d.id = ?`
	rows, err := d.db.Query(statement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	destinations := []Destination{}
	for rows.Next() {
		var destination Destination
		err = rows.Scan(&destination.ID, &destination.Name, &destination.Location, &destination.Description, &destination.BudayaName, &destination.Photo_path)
		if err != nil {
			return nil, err
		}
		destinations = append(destinations, destination)
	}
	return destinations, nil
}

//function to create destination including review, budaya, and photos entities
func (d *DestinationModels) CreateDestination(destination *Destination) error {
	statement := `INSERT INTO destinations (name, location, description) VALUES (?, ?, ?)`
	result, err := d.db.Exec(statement, destination.Name, destination.Location, destination.Description)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	destination.ID = int(id)
	return nil
}

//function to search destination by name including review, budaya, and photos entities
func (d *DestinationModels) SearchDestinationbyName(name string) ([]Destination, error) {
	statement := `SELECT 
	d.*, b.name, p.path
	 FROM destinations d
	 JOIN budaya b ON d.id = b.destination_id
	 JOIN photos p ON d.id = p.destination_id
	 WHERE d.name LIKE ?`
	rows, err := d.db.Query(statement, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	destinations := []Destination{}
	for rows.Next() {
		var destination Destination
		err = rows.Scan(&destination.ID, &destination.Name, &destination.Location, &destination.Description, &destination.BudayaName, &destination.Photo_path)
		if err != nil {
			return nil, err
		}
		destinations = append(destinations, destination)
	}
	return destinations, nil
}

//function to search destination by location including review, budaya, and photos entities
func (d *DestinationModels) SearchDestinationbyLocation(location string) ([]Destination, error) {
	statement := `SELECT 
	d.*, b.name, p.path
	 FROM destinations d
	 JOIN budaya b ON d.id = b.destination_id
	 JOIN photos p ON d.id = p.destination_id
	 WHERE d.location LIKE ?`
	rows, err := d.db.Query(statement, "%"+location+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	destinations := []Destination{}
	for rows.Next() {
		var destination Destination
		err = rows.Scan(&destination.ID, &destination.Name, &destination.Location, &destination.Description, &destination.BudayaName, &destination.Photo_path)
		if err != nil {
			return nil, err
		}
		destinations = append(destinations, destination)
	}
	return destinations, nil
}

//function to delete destination
func (d *DestinationModels) DeleteDestination(id int) error {
	statement := `DELETE FROM destinations WHERE id = ?`
	_, err := d.db.Exec(statement, id)
	if err != nil {
		return err
	}
	return nil
}
