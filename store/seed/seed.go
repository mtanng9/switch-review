package seed

import (
	"database/sql"
	"log"
	"math/rand/v2"
)

func SeedDB(db *sql.DB) {
	err := seedSwitchOneReviews(db)
	if err != nil {
		log.Fatalf("Failed to seed Switch 1 Reviews data: %v", err)
	}

	err = seedSwitchTwoReviews(db)
	if err != nil {
		log.Fatalf("Failed to seed Switch 2 Reviews data: %v", err)
	}

	err = seedArticles(db)
	if err != nil {
		log.Fatalf("Failed to seed Articles data: %v", err)
	}
}

func seedSwitchOneReviews(db *sql.DB) error {
	for _, switchOneReview := range switchOneReviewTitles {
		createReview := `
		INSERT INTO reviews (title, snippet, device, score, body)
		VALUES (?, ?, ?, ?, ?);
		`

		_, err := db.Exec(createReview, switchOneReview, loremIpsumSnippet, "Switch 1", randomScore(), loremIpsumBody)
		if err != nil {
			return err
		}
	}
	return nil
}

func seedSwitchTwoReviews(db *sql.DB) error {
	for _, switchTwoReview := range switchTwoReviewTitles {
		createReview := `
		INSERT INTO reviews (title, snippet, device, score, body)
		VALUES 
		(?, ?, ?, ?, ?);
		`

		_, err := db.Exec(createReview, switchTwoReview, loremIpsumSnippet, "Switch 2", randomScore(), loremIpsumBody)
		if err != nil {
			return err
		}
	}
	return nil
}

func seedArticles(db *sql.DB) error {
	for _, articleTitle := range articleTitles {
		createArticle := `
			INSERT INTO articles (title, snippet, body)
			VALUES 
			(?, ?, ?);
			`

		_, err := db.Exec(createArticle, articleTitle, loremIpsumSnippet, loremIpsumBody)
		if err != nil {
			return err
		}
	}
	return nil
}

func randomScore() int {
	min := 85
	max := 100

	return rand.IntN(max-min+1) + min
}
