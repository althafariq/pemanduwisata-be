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

		CREATE TABLE IF NOT EXISTS destinations (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) NOT NULL,
			location VARCHAR(100) NOT NULL,
			description TEXT NOT NULL,
			budaya_name VARCHAR(100) NULL,
			budaya_description TEXT NULL,
			photo_path VARCHAR(100) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS reviews (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			destination_id INTEGER NOT NULL,
			rating INTEGER NOT NULL,
			review TEXT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES user(user_id),
			FOREIGN KEY (destination_id) REFERENCES destinations(id)
		);

		CREATE TABLE IF NOT EXISTS TelpDarurat (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) NOT NULL,
			number VARCHAR(100) NOT NULL
		);
	`)

	if err != nil {
		panic(err)
	}
	
	seeder.Seed(db)
}