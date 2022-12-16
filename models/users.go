package models

import (
	"database/sql"
	"errors"
	"net/http"
	"regexp"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type UserModels struct {
	db *sql.DB
}

func NewUserModels(db *sql.DB) *UserModels {
	return &UserModels{
		db: db,
	}
}

func (u *UserModels) GetUserRole(id int) (*string, error) {
	statement := "SELECT role FROM user WHERE user_id = ?"
	var role string
	res := u.db.QueryRow(statement, id)
	err := res.Scan(&role)
	return &role, err
}

func (u *UserModels) GetUserData(id int) (*User, error) {
	statement := "SELECT * FROM user WHERE user_id = ?"
	var user User
	res := u.db.QueryRow(statement, id)
	err := res.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.Profile_pic, &user.Role, &user.CreatedAt)
	return &user, err
}

func (u *UserModels) Login(email string, password string) (*int, error) {
	sql := `SELECT user_id, password FROM user WHERE email = ?`
	res := u.db.QueryRow(sql, email, password)

	var hashedPassword string
	var id int

	res.Scan(&id, &hashedPassword)

	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return nil, errors.New("Login Failed")
	}

	return &id, nil
}

func (u *UserModels) CheckEmail(email string) (bool, error) {
	statement := "SELECT count(*) FROM user WHERE email = ?"
	res := u.db.QueryRow(statement, email)

	var count int
	err := res.Scan(&count)
	if count > 0 {
		return false, err
	}
	return true, err
} 

func (u *UserModels) Register(firstname string, lastname string, email string, password string) (userId int, responseCode int, err error) {
	uniqueEmail, err := u.CheckEmail(email)
	if err != nil {
		return -1, http.StatusInternalServerError, err
	}

	if !uniqueEmail {
		return -1, http.StatusBadRequest, errors.New("email already registered")
	}

	formatEmail, err := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if err != nil {
		return -1, http.StatusInternalServerError, err
	}

	isValid := formatEmail.Match([]byte(email))
	if !isValid {
		return -1, http.StatusBadRequest, errors.New("invalid email format")
	}

	minPassword, err := regexp.Compile(`^[a-zA-Z0-9]{8,}$`)
	if err != nil {
		return -1, http.StatusInternalServerError, err
	}

	checkPassword := minPassword.Match([]byte(password))
	if !checkPassword {
		return -1, http.StatusBadRequest, errors.New("password must be at least 8 characters")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	statement := "INSERT INTO user (firstname, lastname, email, password, role) VALUES (?, ?, ?, ?, ?)"

	res, err := u.db.Exec(statement, firstname, lastname, email, hashedPassword, "user")
	if err != nil {
		return -1, http.StatusInternalServerError, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return -1, http.StatusInternalServerError, err
	}

	return int(lastId), http.StatusOK, err
}

func (u *UserModels) UpdateUserData(id int, firstname, lastname, email string) error {
	user, err := u.GetUserData(id)
	if err != nil {
		return err
	}

	if user.Email != email {
		uniqueEmail, err := u.CheckEmail(email)

		if err != nil {
			return err
		}

		regex, err := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
		if err != nil {
			return err
		}

		isValid := regex.Match([]byte(email))
		if !isValid {
			return errors.New("invalid email")
		}

		if !uniqueEmail {
			return errors.New("email already registered")
		}
	}

	statement := "UPDATE user SET firstname = ?, lastname = ?, email = ? WHERE user_id = ?"

	_, err = u.db.Exec(statement, firstname, lastname, email, id)
	return err
}

func (u *UserModels) UpdateAvatar(userId int, filepath string) error {
	statement := "UPDATE user SET profile_pic = ? WHERE user_id = ?"
	_, err := u.db.Exec(statement, filepath, userId)
	return err
}