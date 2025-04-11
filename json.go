package main

import (
	"encoding/json" // For JSON encoding/decoding
	"log"           // For logging
	"net/http"      // For HTTP functionality
)

// respondWithError sends an error response in JSON format
// Parameters:
//   - w: HTTP response writer
//   - code: HTTP status code (e.g., 400 for bad request, 500 for server error)
//   - message: Error message to be sent to the client
func respondWithError(w http.ResponseWriter, code int, message string) {
	// Log 5XX errors for debugging purposes
	if code > 499 {
		log.Printf("Responding with 5XX error: %v", message)
	}

	// Define the structure of error responses
	// This ensures consistent error response format
	type errorResponse struct {
		Error string `json:"error"` // The error message field in JSON response
	}

	// Send the error response using the JSON helper
	// This will set the appropriate status code and content type
	respondWithJSON(w, code, errorResponse{
		Error: message,
	})
}

// respondWithJSON sends a JSON response with the given status code and payload
// Parameters:
//   - w: HTTP response writer
//   - code: HTTP status code (e.g., 200 for success)
//   - payload: Any data that can be converted to JSON
func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	// Convert the payload to JSON bytes
	dat, err := json.Marshal(payload)
	if err != nil {
		// If JSON marshaling fails, log the error and return 500
		log.Printf("Failed to marshal JSON response: %v", err)
		w.WriteHeader(500)
		return
	}

	// Set the content type header to application/json
	// This tells the client to expect JSON data
	w.Header().Add("Content-Type", "application/json")

	// Set the HTTP status code
	// This indicates the success or failure of the request
	w.WriteHeader(code)

	// Write the JSON response to the client
	w.Write(dat)
}
