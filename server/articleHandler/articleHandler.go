package articlehandler

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/mtanng9/switch-review/store/article"
)

type articlesPage struct {
	Title    string
	Articles []article.Article
}

func GetArticle(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		baseFile := path.Join("html", "base.html")

		article, err := article.GetArticleById(db, id)
		if err != nil {
			errorFile := path.Join("html", "error.html")
			t, _ := template.ParseFiles(baseFile, errorFile)
			t.Execute(w, nil)
			return
		}

		articleFile := path.Join("html", "article.html")
		t, _ := template.ParseFiles(baseFile, articleFile)
		t.Execute(w, article)

	}
}

func GetAllArticles(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		baseFile := path.Join("html", "base.html")

		var articles []article.Article
		articles, err := article.GetAllArticles(db)
		if err != nil {
			fmt.Printf("Could not get Articles: %v \n", err)
			errorFile := path.Join("html", "error.html")
			t, _ := template.ParseFiles(baseFile, errorFile)
			t.Execute(w, nil)
			return
		}

		articlesPage := articlesPage{
			Title:    "Articles",
			Articles: articles,
		}

		articlesFile := path.Join("html", "articles.html")
		t, _ := template.ParseFiles(baseFile, articlesFile)
		t.Execute(w, articlesPage)
	}
}
