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

type Review struct {
	ID 									int    		`json:"id"`
	UserID 							int    		`json:"user_id"`
	DestinationID 			int    		`json:"destination_id"`
	Rating 							int    		`json:"rating"`
	Review 							string 		`json:"review"`
	CreatedAt 					time.Time `json:"created_at"`
	Firstname 					string 		`json:"firstname"`
	Lastname 						string 		`json:"lastname"`
	Profile_pic 				*string 	`json:"profile_pic"`
	DestinationName 		string 		`json:"name"`
	DestinationLocation string 		`json:"location"`
}

type TelpDarurat struct {
	ID 				int 			`json:"id"`
	Name 			string 		`json:"name"`
	Number		string 		`json:"number"`
}

type Destination struct {
	ID 									int 			`json:"id"`
	Name 								string 		`json:"name"`
	Location 						string 		`json:"location"`
	Description 				string 		`json:"description"`
	BudayaName 					string 		`json:"budaya_name"`
	Photo_path 					string 		`json:"photo_path"`

}

type Budaya struct {
	ID 									int 			`json:"id"`
	DestinationID 			int 			`json:"destination_id"`
	Name 								string 		`json:"name"`
	Description 				string 		`json:"description"`
}

type Photos struct {
	photo_id 				int 			`json:"photo_id"`
	destination_id 	int		 		`json:"destination_id"`
	path 						string 		`json:"path"`
}