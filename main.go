package main

import (
	"log"
	"net/http"
)

var transactions []Transaction // Определение среза Transaction

func main() {
	r := NewRouter()
	log.Println("Server is starting on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
