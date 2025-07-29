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

func SeedDB(db *sql.DB) {
	createReview := `
		INSERT INTO reviews (title, snippet, device, score, body)
		VALUES 
		('Review 1', 'Review 1 Snippet', 'Switch 1', 85, 'Review 1 Body'),
		('Review 2', 'Review 2 Snippet', 'Switch 1', 85, 'Review 2 Body'),
		('Review 3', 'Review 3 Snippet', 'Switch 1', 85, 'Review 3 Body'),
		('Review 4', 'Review 4 Snippet', 'Switch 1', 85, 'Review 4 Body'),
		('Review 5', 'Review 5 Snippet', 'Switch 1', 85, 'Review 5 Body');
	`

	_, err := db.Exec(createReview)
	if err != nil {
		log.Fatalf("Failed to seed Reviews data: %v", err)
	}

	createArticle := `
		INSERT INTO articles (title, snippet, body)
		VALUES 
		('Article 1', 'Article 1 Snippet', 'Article 1 Body'),
		('Article 2', 'Article 2 Snippet', 'Article 2 Body'),
		('Article 3', 'Article 3 Snippet', 'Article 3 Body'),
		('Article 4', 'Article 4 Snippet', 'Article 4 Body');
	`

	_, err = db.Exec(createArticle)
	if err != nil {
		log.Fatalf("Failed to seed Articles data: %v", err)
	}
}
