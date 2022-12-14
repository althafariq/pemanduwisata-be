package main

import (
	"database/sql"

	"github.com/althafariq/pemanduwisata-be/db/migration"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./pemandu.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migration.Migrate(db)
}