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

func GetArticleById(db *sql.DB, id string) (Article, error) {
	var article Article
	query := `SELECT id, title, snippet, body, created_at FROM articles WHERE id = ?`
	err := db.QueryRow(query, id).Scan(&article.Id, &article.Title, &article.Snippet, &article.Body, &article.CreatedAt)
	if err != nil {
		return article, err
	}

	return article, nil
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

func GetAllArticles(db *sql.DB) ([]Article, error) {
	var articles []Article
	query := `SELECT id, title, snippet, body, created_at FROM articles ORDER BY created_at`
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
