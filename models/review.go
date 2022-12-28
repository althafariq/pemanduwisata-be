package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)


type ReviewModels struct {
	db *sql.DB
}

func NewReviewModels(db *sql.DB) *ReviewModels {
	return &ReviewModels{
		db: db,
	}
}

func (r *ReviewModels) GetReviewbyDestinationID(destinationID int) ([]Review, error) {
	statement := `SELECT 
	r.*, u.firstname, u.lastname, u.profile_pic
	 FROM reviews r
	 LEFT JOIN user u ON r.user_id = u.user_id
	 LEFT JOIN destinations d ON r.destination_id = d.id
	 WHERE r.destination_id = ?
	 ORDER BY r.created_at DESC`
	rows, err := r.db.Query(statement, destinationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviews := []Review{}
	for rows.Next() {
		var review Review
		err = rows.Scan(
			&review.ID, 
			&review.UserID, 
			&review.DestinationID, 
			&review.Rating, 
			&review.Review, 
			&review.CreatedAt, 
			&review.Firstname, 
			&review.Lastname, 
			&review.Profile_pic,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

func (r *ReviewModels) GetReviewbyUserID(reviewID int) (int, error) {
	statement := `SELECT user_id FROM reviews WHERE id = ?`
	
	var userID int
	err := r.db.QueryRow(statement, reviewID).Scan(&userID)
	switch err {
	case sql.ErrNoRows:
		return 0, nil
	case nil:
		return userID, nil
	default:
		return 0, err
	} 
}


func (r *ReviewModels) CreateReview(review Review) (int, error) {
	statement := `INSERT INTO reviews (user_id, destination_id, rating, review, created_at) VALUES (?, ?, ?, ?, DATETIME('now'))`
	result, err := r.db.Exec(statement, review.UserID, review.DestinationID, review.Rating, review.Review, review.CreatedAt)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (r *ReviewModels) UpdateReview(review Review) error {
	statement := `UPDATE reviews SET rating = ?, review = ? WHERE id = ?`
	_, err := r.db.Exec(statement, review.Rating, review.Review, review.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReviewModels) DeleteReview(id int) error {
	statement := `DELETE FROM reviews WHERE id = ?`
	_, err := r.db.Exec(statement, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReviewModels) GetAverageRating(id int) (float64, error) {
	statement := `SELECT AVG(rating) FROM reviews WHERE destination_id = ?`
	var avgRating sql.NullFloat64
	err := r.db.QueryRow(statement, id).Scan(&avgRating) 
	if err != nil {
		return -1, err
	}
	return avgRating.Float64, nil
}

func (r *ReviewModels) GetReviewCount(id int) (int, error) {
	statement := `SELECT COUNT(*) FROM reviews WHERE destination_id = ?`
	var reviewCount int
	err := r.db.QueryRow(statement, id).Scan(&reviewCount) 
	if err != nil {
		return -1, err
	}
	return reviewCount, nil
}