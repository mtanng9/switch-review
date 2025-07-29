package review

import (
	"database/sql"
	"time"
)

type Review struct {
	Id        int
	Title     string
	Snippet   string
	Device    string
	Score     int
	Body      string
	createdAt time.Time
}

func GetReviewById(db *sql.DB, id string) (Review, error) {
	var review Review
	query := `SELECT id, title, snippet, device, score, body, created_at FROM reviews WHERE id = ?`
	err := db.QueryRow(query, id).Scan(&review.Id, &review.Title, &review.Snippet, &review.Device, &review.Score, &review.Body, &review.createdAt)
	if err != nil {
		return review, err
	}
	return review, nil
}

func GetReviewsByDevice(db *sql.DB, device string) ([]Review, error) {
	var reviews []Review
	query := `SELECT id, title, snippet, device, score, body, created_at FROM reviews WHERE device = ?`
	rows, err := db.Query(query, device)
	if err != nil {
		return reviews, err
	}

	return scanReviews(rows)
}

func GetFourLatestReviews(db *sql.DB) ([]Review, error) {
	var reviews []Review
	query := `SELECT id, title, snippet, device, score, body, created_at FROM reviews ORDER BY created_at DESC LIMIT 4`
	rows, err := db.Query(query)
	if err != nil {
		return reviews, err
	}

	return scanReviews(rows)
}

func scanReviews(rows *sql.Rows) ([]Review, error) {
	var reviews []Review

	for rows.Next() {
		var review Review
		err := rows.Scan(&review.Id, &review.Title, &review.Snippet, &review.Device, &review.Score, &review.Body, &review.createdAt)
		if err != nil {
			return reviews, err
		}
		reviews = append(reviews, review)
	}

	if err := rows.Err(); err != nil {
		return reviews, err
	}

	return reviews, nil
}
