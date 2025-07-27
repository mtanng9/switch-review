package reviewhandler

import (
	"database/sql"
	"log"
	"net/http"
	"path"
	"text/template"
	"time"
)

type Review struct {
	Id        int
	Title     string
	Snippet   string
	Score     int
	Body      string
	createdAt time.Time
}

// CRUD
// CREATE
// READ
// UPDATE
// DESTROY

func GetReview(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		var review Review
		query := `SELECT id, title, snippet, score, body, created_at FROM reviews WHERE id = ?`
		err := db.QueryRow(query, id).Scan(&review.Id, &review.Title, &review.Snippet, &review.Score, &review.Body, &review.createdAt)
		if err != nil {
			log.Fatalf("Could not get Review: %v", err)
		}

		baseFile := path.Join("html", "base.html")
		reviewFile := path.Join("html", "review.html")
		t, _ := template.ParseFiles(baseFile, reviewFile)
		t.Execute(w, review)
	}
}
