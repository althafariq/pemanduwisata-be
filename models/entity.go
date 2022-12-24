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
	DestinationID 			int    		`json:"destination_id"`
	UserID 							int    		`json:"user_id"`
	Firstname 					string 		`json:"firstname"`
	Lastname 						string 		`json:"lastname"`
	Profile_pic 				*string 	`json:"profile_pic"`
	Rating 							int    		`json:"rating"`
	Review 							string 		`json:"review"`
	CreatedAt 					time.Time `json:"created_at"`
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
	BudayaDescription 	string 		`json:"budaya_description"`
	PhotoPath 					string 		`json:"photo_path"`
	AvgRating 					float64 	`json:"avg_rating"`
	TotalReview 				int 			`json:"total_review"`
}