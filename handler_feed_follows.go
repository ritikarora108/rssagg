package main

import (
	"encoding/json" // For JSON encoding/decoding
	"fmt"           // For formatted strings
	"net/http"      // For HTTP functionality
	"time"          // For time operations

	"github.com/go-chi/chi/v5"                          // For URL parameter extraction
	"github.com/google/uuid"                            // For UUID generation
	"github.com/ritikarora108/rssagg/internal/database" // Our database package
)

// HandlerCreateFeedFollow handles the creation of a new feed follow relationship
// It expects a POST request with a JSON body containing a feed_id
// The user is authenticated via middleware and passed as a parameter
func (apiCfg *apiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	// Define the structure of the expected JSON request body
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"` // The feed ID to follow
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

	// Create a new feed follow record in the database
	feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),       // Generate a new UUID for the feed follow
		CreatedAt: time.Now().UTC(), // Set creation time to current UTC time
		UpdatedAt: time.Now().UTC(), // Set update time to current UTC time
		UserID:    user.ID,          // Use the authenticated user's ID
		FeedID:    params.FeedID,    // Use the feed ID from the request
	})
	if err != nil {
		// If database operation fails, return a 400 error
		respondWithError(w, 400, fmt.Sprintf("Error creating feed follow: %v", err))
		return
	}

	// If everything succeeds, return the created feed follow with 201 status code
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feed))
}

// HandlerGetFeedFollowsByUser retrieves all feeds that a user is following
// The user is authenticated via middleware and passed as a parameter
func (apiCfg *apiConfig) HandlerGetFeedFollowsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// Get all feed follows for the authenticated user
	feeds, err := apiCfg.DB.GetFeedFollowsByUser(r.Context(), user.ID)
	if err != nil {
		// If database operation fails, return a 400 error
		respondWithError(w, 400, fmt.Sprintf("Error getting feeds: %v", err))
		return
	}

	// Return the list of feed follows with 200 status code
	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feeds))
}

// HandlerDeleteFeedFollow deletes a feed follow relationship
// It expects a DELETE request with the feed follow ID in the URL
// The user is authenticated via middleware and passed as a parameter
func (apiCfg *apiConfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	// Get the feed follow ID from the URL parameters
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")

	// Parse the string ID into a UUID
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		// If ID parsing fails, return a 400 error
		respondWithError(w, 400, fmt.Sprintf("Error parsing feed follow ID: %v", err))
		return
	}

	// Delete the feed follow from the database
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID, // The feed follow ID to delete
		UserID: user.ID,      // The authenticated user's ID
	})
	if err != nil {
		// If database operation fails, return a 400 error
		respondWithError(w, 400, fmt.Sprintf("Error deleting feed follow: %v", err))
		return
	}

	// Verify that the feed follow was deleted and belonged to the user
	feedFollow, err := apiCfg.DB.GetFeedFollow(r.Context(), feedFollowID)
	if err == nil && feedFollow.UserID != user.ID {
		// If the feed follow exists and doesn't belong to the user, return a 400 error
		respondWithError(w, 400, fmt.Sprintln("User is not authorized to delete this feed follow"))
		return
	}

	// Return success message with 200 status code
	respondWithJSON(w, 200, map[string]string{"message": "Feed follow deleted"})
}
