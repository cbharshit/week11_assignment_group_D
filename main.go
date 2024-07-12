package main

import (
	"fmt"
	"log"
	"net/http"
)

// Function for Home Handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}
	fmt.Fprintln(w, "Welcome to Home!")
}

// Function for About handler
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to About page.")
}

// Function for Product handler
func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Product page.")
}

func main() {
	mux := http.NewServeMux()

	// Register the handlers
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/product", productHandler)

	log.Println("Starting server on :3001")

	// Starting the HTTP server on port 3001
	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
