package indexhandler

import (
	"database/sql"
	"html/template"
	"net/http"
	"path"

	"github.com/mtanng9/switch-review/store/article"
	"github.com/mtanng9/switch-review/store/review"
)

type IndexData struct {
	Reviews  []review.Review
	Articles []article.Article
}

func GetIndex(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		baseFile := path.Join("html", "base.html")
		indexFile := path.Join("html", "index.html")

		reviews, err := review.GetFourLatestReviews(db)
		if err != nil {
			notFoundFile := path.Join("html", "error.html")
			t, _ := template.ParseFiles(baseFile, notFoundFile)
			t.Execute(w, nil)
			return
		}

		articles, err := article.GetThreeLatestArticles(db)
		if err != nil {
			notFoundFile := path.Join("html", "error.html")
			t, _ := template.ParseFiles(baseFile, notFoundFile)
			t.Execute(w, nil)
			return
		}

		indexInput := IndexData{
			Reviews:  reviews,
			Articles: articles,
		}

		t, _ := template.ParseFiles(baseFile, indexFile)
		t.Execute(w, indexInput)
	}
}
