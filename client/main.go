package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cb", callbackHandler)

	// Start the server on port 8089
	fmt.Println("Listening on port 8089...")
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// callbackHandler handles requests to the /cb endpoint
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request on /cb")
	log.Printf("Method: %s, URL: %s", r.Method, r.URL.String())

	// Log query parameters
	for key, values := range r.URL.Query() {
		for _, value := range values {
			log.Printf("Query Param: %s = %s", key, value)
		}
	}

	code := r.URL.Query().Get("code")
	// Respond with a simple message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Callback received " + code))
}
