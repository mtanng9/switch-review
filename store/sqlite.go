package store

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./switch-review-website.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	log.Println("Connected to database successfully")

	createReviewsTable := `
	CREATE TABLE IF NOT EXISTS reviews (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		snippet TEXT NOT NULL,
		device TEXT NOT NULL,
		score INTEGER NOT NULL,
		body TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err = db.Exec(createReviewsTable)
	if err != nil {
		log.Fatalf("Failed to create Reviews table: %v", err)
	}

	createArticleTable := `
	CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		snippet TEXT NOT NULL,
		body TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`
	_, err = db.Exec(createArticleTable)
	if err != nil {
		log.Fatalf("Failed to create Articles table: %v", err)
	}

	return db
}
