package main

import (
	"time"

	"github.com/google/uuid"                            // For UUID handling
	"github.com/ritikarora108/rssagg/internal/database" // Our database package
)

// User represents the user model in our API responses
// This is the structure that clients will receive
type User struct {
	ID        uuid.UUID `json:"id"`         // Unique identifier for the user
	CreatedAt time.Time `json:"created_at"` // When the user was created
	UpdatedAt time.Time `json:"updated_at"` // When the user was last updated
	Name      string    `json:"name"`       // User's name
	ApiKey    string    `json:"api_key"`    // User's API key
}

// databaseUserToUser converts a database user model to our API user model
// This function helps us maintain a clean separation between our database
// models and our API models
func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,        // Copy the ID
		CreatedAt: dbUser.CreatedAt, // Copy the creation time
		UpdatedAt: dbUser.UpdatedAt, // Copy the update time
		Name:      dbUser.Name,      // Copy the name
		ApiKey:    dbUser.ApiKey,    // Copy the API key
	}
}


type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,	
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}


type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,	
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFollows
}









