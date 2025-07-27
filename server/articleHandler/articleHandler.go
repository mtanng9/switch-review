package articlehandler

import (
	"database/sql"
	"html/template"
	"log"
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

		var article Article
		query := `SELECT id, title, snippet, body, created_at FROM articles WHERE id = ?`
		err := db.QueryRow(query, id).Scan(&article.Id, &article.Title, &article.Snippet, &article.Body, &article.CreatedAt)
		if err != nil {
			log.Fatalf("Could not get Article: %v", err)
		}

		baseFile := path.Join("html", "base.html")
		articleFile := path.Join("html", "article.html")
		t, _ := template.ParseFiles(baseFile, articleFile)
		t.Execute(w, article)

	}
}
