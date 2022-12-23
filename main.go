package main

import (
	"database/sql"

	"github.com/althafariq/pemanduwisata-be/controllers"
	"github.com/althafariq/pemanduwisata-be/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./pemandu.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	usersModel := models.NewUserModels(db)
	destinationModel := models.NewDestinationModels(db)
	reviewModel := models.NewReviewModels(db)
	// budayaModel := models.NewBudayaModels(db)
	telpDaruratModel := models.NewTelpDaruratModels(db)

	mainApi := controllers.NewApi(*usersModel, *destinationModel, *reviewModel, *telpDaruratModel)
	mainApi.Start()
}