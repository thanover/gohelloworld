package main

import (
	"fmt"
	"net/http"
)



func main() {
	// Create a new HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Handle the request
			fmt.Fprintf(w, "Hello, World!")
		}),
	}

	// Start the server
	fmt.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}