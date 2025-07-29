package server

import (
	"database/sql"
	"log"
	"net/http"

	articlehandler "github.com/mtanng9/switch-review/server/articleHandler"
	indexhandler "github.com/mtanng9/switch-review/server/indexHandler"
	reviewhandler "github.com/mtanng9/switch-review/server/reviewHandler"
)

// Start Server is a function that take a pointer to a sqlite 3 DB connection. This function then creates a http.ServeMux
// which is a http server. Then HandlerFuncs are added to the ServeMux to capture traffic to defined path with certain HTTP Methods
// ex. GET /review. Then we log for viability purposes that the server is starting and on what port. Then we start the server in log.Fatal
// so if something error out the server the program exits and logs the error
func StartServer(db *sql.DB) {
	// Removed http for mux to handle path params easier
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	mux.Handle("GET /css/", http.StripPrefix("/css/", http.FileServer(http.Dir("web/css"))))

	mux.HandleFunc("GET /", indexhandler.GetIndex(db))
	mux.HandleFunc("GET /review/switch1", reviewhandler.GetReviewByDevice(db, "Switch 1"))
	mux.HandleFunc("GET /review/{id}", reviewhandler.GetReview(db))
	mux.HandleFunc("GET /article/{id}", articlehandler.GetArticle(db))

	log.Println("Starting the server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
