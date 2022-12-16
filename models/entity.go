package models

import "time"

type User struct {
	ID 					int 			`json:"id"`
	Firstname 	string 		`json:"firstname"`
	Lastname 		string 		`json:"lastname"`
	Email 			string 		`json:"email"`
	Password 		string 		`json:"password"`
	Profile_pic *string 	`json:"profile_pic"`
	Role 				string 		`json:"role"`
	CreatedAt 	time.Time `json:"created_at"`
}

type review struct {
	id 							int 			`db:"id"`
	user_id 				int 			`db:"user_id"`
	destination_id 	int 			`db:"destination_id"`
	rating 					int 			`db:"rating"`
	createdAt 			time.Time `db:"created_at"`
}

type telpDarurat struct {
	id 				int 			`db:"id"`
	name 			string 		`db:"name"`
	phone			string 		`db:"phone"`
}

type Destination struct {
	id 				int				`db:"destination_id"`
	name 			string		`db:"name"`
	location 	string		`db:"location"`
}

type Budaya struct {
	budaya_id 			int 			`db:"budaya_id"`
	destination_id 	int 			`db:"destination_id"`
	name 						string 		`db:"name"`
	description 		string 		`db:"description"`
}