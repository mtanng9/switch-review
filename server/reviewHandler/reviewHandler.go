package reviewhandler

import (
	"database/sql"
	"net/http"
	"path"
	"text/template"

	"github.com/mtanng9/switch-review/store/review"
)

type ReviewByDevicePage struct {
	Title   string
	Reviews []review.Review
}

// CRUD
// CREATE
// READ
// UPDATE
// DESTROY

func GetReview(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		baseFile := path.Join("html", "base.html")

		review, err := review.GetReviewById(db, id)
		if err != nil {
			notFoundFile := path.Join("html", "notFound.html")
			t, _ := template.ParseFiles(baseFile, notFoundFile)
			t.Execute(w, nil)
			return
		}

		reviewFile := path.Join("html", "review.html")
		t, _ := template.ParseFiles(baseFile, reviewFile)
		t.Execute(w, review)
	}
}

func GetReviewByDevice(db *sql.DB, device string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		baseFile := path.Join("html", "base.html")

		reviews, err := review.GetReviewsByDevice(db, device)
		if err != nil {
			notFoundFile := path.Join("html", "error.html")
			t, _ := template.ParseFiles(baseFile, notFoundFile)
			t.Execute(w, nil)
			return
		}

		reviewByDevicePage := ReviewByDevicePage{
			Title:   device,
			Reviews: reviews,
		}
		reviewFile := path.Join("html", "reviewsByDevice.html")
		t, _ := template.ParseFiles(baseFile, reviewFile)
		t.Execute(w, reviewByDevicePage)
	}
}
