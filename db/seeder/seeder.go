package seeder

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func Seed(db *sql.DB) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("tgadmin123"), bcrypt.DefaultCost)

	admin, err := db.Exec("INSERT INTO user (firstname, lastname, email, password, role) VALUES ('Admin', 'Tour Guide', 'admin@email.com', ?, 'admin')", hashedPassword)
	if err != nil {
		panic(err)
	}
	_, err = admin.LastInsertId()
	if err != nil {
		panic(err)
	}
}
