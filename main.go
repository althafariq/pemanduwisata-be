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

	mainApi := controllers.NewApi(*usersModel)
	mainApi.Start()
}