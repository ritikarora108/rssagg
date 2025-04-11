package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ritikarora108/rssagg/internal/database"
)




func startScrapping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequests time.Duration,
) {
	log.Printf("Scrapping on %v goroutines every %s duration", concurrency, timeBetweenRequests)


	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Printf("Error fetching feeds: %v", err)
			continue
		}
		
		// for _, feed := range feeds {
		// 	log.Printf("Scrapping feed %v", feed.ID)
		// 	db.MarkFeedAsFetched(
		// 		context.Background(),
		// 		feed.ID,
		// 	)
		// }

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db,wg,feed)
		}
		wg.Wait()
	}
}

func scrapeFeed( db *database.Queries , wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	
	log.Printf("Scrapping feed %v", feed.ID)
	_,err := db.MarkFeedAsFetched(
		context.Background(),
		feed.ID,
	)

	if err != nil {
		log.Printf("Error marking feed as fetched: %v", err)
		return
	}


	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed for %v: %v", feed.Url, err)
		return
	}

	for _, item := range rssFeed.Channel.Item {


		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Couldn't parse date %v with err: %v", item.PubDate, err)
			continue
		}

		_, err = db.CreatePost(context.Background(),database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid: item.Description != "",
			},
			PublishedAt: publishedAt,
			Url: item.Link,
			FeedID: feed.ID,
		})
		if err != nil {
			if(!strings.Contains(err.Error(),"duplicate key value violates unique constraint")) {
				log.Printf("Couldn't create post for %v: %v", item.Title, err)
			}
			continue
		}
	}

	log.Printf("Feed %v has %v posts", feed.ID, len(rssFeed.Channel.Item))


}
