package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API key from the request headers
// Returns the API key if found, otherwise returns an error

// Example :
// Authorization: ApiKey {insert api key here}
func GetAPIKey(headers http.Header) (string, error) {
	
	// Extract the API key from the request headers
	apiKey := headers.Get("Authorization")

	// If the API key is not found, return an error
	if apiKey == "" {
		return "", errors.New("No Authentication header found")
	}


	// Remove the "ApiKey " prefix from the API key
	apiKeyPieces := strings.Split(apiKey, " ")
	
	if len(apiKeyPieces) != 2 || apiKeyPieces[0] != "ApiKey" {
		return "", errors.New("Invalid Authentication header")
	}

	// Return the API key
	return apiKeyPieces[1], nil
}
