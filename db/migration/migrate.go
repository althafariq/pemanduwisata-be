package migration

import (
	"database/sql"

	"github.com/althafariq/pemanduwisata-be/db/seeder"
	_ "github.com/mattn/go-sqlite3"
)

func Migrate(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS user (
			user_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			firstname VARCHAR(50) NOT NULL,
			lastname VARCHAR(50) NOT NULL,
			email VARCHAR(100) NOT NULL,
			password VARCHAR(300) NOT NULL,
			profile_pic VARCHAR(100) DEFAULT NULL,
			role VARCHAR(10) NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)

	if err != nil {
		panic(err)
	}
	
	seeder.Seed(db)
}