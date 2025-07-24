package server

import (
	"log"
	"net/http"
)

func StartServer() {
	log.Println("Starting the server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
