package migration

import (
	"database/sql"

	"github.com/althafariq/pemanduwisata-be/db/seeder"
	_ "github.com/mattn/go-sqlite3"
)

func Migrate(db *sql.DB) {
	seeder.Seed(db)
}