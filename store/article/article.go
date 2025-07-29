package article

import (
	"database/sql"
	"time"
)

type Article struct {
	Id        int
	Title     string
	Snippet   string
	Body      string
	CreatedAt time.Time
}

func GetThreeLatestArticles(db *sql.DB) ([]Article, error) {
	var articles []Article
	query := `SELECT id, title, snippet, body, created_at FROM articles ORDER BY created_at DESC LIMIT 3`
	rows, err := db.Query(query)
	if err != nil {
		return articles, err
	}

	return scanArticles(rows)
}

func scanArticles(rows *sql.Rows) ([]Article, error) {
	var articles []Article

	for rows.Next() {
		var article Article
		err := rows.Scan(&article.Id, &article.Title, &article.Snippet, &article.Body, &article.CreatedAt)
		if err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		return articles, err
	}

	return articles, nil
}
