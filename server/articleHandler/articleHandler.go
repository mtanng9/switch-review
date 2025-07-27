package articlehandler

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"time"
)

type Article struct {
	Id        int
	Title     string
	Snippet   string
	Body      string
	CreatedAt time.Time
}

func GetArticle(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		baseFile := path.Join("html", "base.html")

		var article Article
		query := `SELECT id, title, snippet, body, created_at FROM articles WHERE id = ?`
		err := db.QueryRow(query, id).Scan(&article.Id, &article.Title, &article.Snippet, &article.Body, &article.CreatedAt)
		if err != nil {
			fmt.Printf("Could not get Article: %v \n", err)
			notFoundFile := path.Join("html", "notFound.html")
			t, _ := template.ParseFiles(baseFile, notFoundFile)
			t.Execute(w, nil)
			return
		}

		articleFile := path.Join("html", "article.html")
		t, _ := template.ParseFiles(baseFile, articleFile)
		t.Execute(w, article)

	}
}
