package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"                            // For generating unique IDs
	"github.com/ritikarora108/rssagg/internal/database" // Our database package
)

// HandlerCreateUser handles the creation of a new user
// It expects a POST request with a JSON body containing a "name" field
// Returns the created user or an error message
func (apiCfg *apiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// Define the structure of the expected JSON request body
	type parameters struct {
		Name string `json:"name"` // The name field in the JSON request
	}

	// Create a JSON decoder that reads from the request body
	decoder := json.NewDecoder(r.Body)

	// Create an empty parameters struct to store the decoded data
	params := parameters{}

	// Try to decode the JSON from request body into our params struct
	err := decoder.Decode(&params)
	if err != nil {
		// If JSON parsing fails, return a 400 error
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Create a new user in the database with the provided name
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),       // Generate a new UUID for the user
		CreatedAt: time.Now().UTC(), // Set creation time to current UTC time
		UpdatedAt: time.Now().UTC(), // Set update time to current UTC time
		Name:      params.Name,      // Use the name from the request
	})
	if err != nil {
		// If database operation fails, return a 400 error
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	// If everything succeeds, return the created user with 200 status code
	respondWithJSON(w, 201, databaseUserToUser(user))
}


func (apiCfg *apiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}



