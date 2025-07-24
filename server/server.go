package server

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is index")
}

func StartServer() {
	http.HandleFunc("/", index)
	log.Println("Starting the server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
