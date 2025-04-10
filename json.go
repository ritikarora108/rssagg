package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// respondWithError sends an error response in JSON format
// code: HTTP status code (e.g., 400 for bad request, 500 for server error)
// message: Error message to be sent to the client
func respondWithError(w http.ResponseWriter, code int, message string) {
	// Log 5XX errors for debugging
	if code > 499 {
		log.Printf("Responding with 5XX error: %v", message)
	}

	// Define the structure of error responses
	type errorResponse struct {
		Error string `json:"error"` // The error message field in JSON response
	}

	// Send the error response using the JSON helper
	respondWithJSON(w, code, errorResponse{
		Error: message,
	})
}

// respondWithJSON sends a JSON response with the given status code and payload
// code: HTTP status code (e.g., 200 for success)
// payload: Any data that can be converted to JSON
func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	// Convert the payload to JSON
	dat, err := json.Marshal(payload)
	if err != nil {
		// If JSON marshaling fails, log the error and return 500
		log.Printf("Failed to marshal JSON response: %v", err)
		w.WriteHeader(500)
		return
	}

	// Set the content type header to application/json
	w.Header().Add("Content-Type", "application/json")

	// Set the HTTP status code
	w.WriteHeader(code)

	// Write the JSON response
	w.Write(dat)
}
