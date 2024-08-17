package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve the password page at the root URL
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./password"))))

	// Start the server on port 8080
	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}