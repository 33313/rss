package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/myshkovsky/rss/internal/database"
)

func StartScraper(db *database.Queries, numFeeds int, interval time.Duration) {
	log.Printf("Scraping %v feeds every %s...", numFeeds, interval)
	ticker := time.NewTicker(interval)

	for ; ; <-ticker.C {
		feedArr, err := db.GetNextFeedsToFetch(context.Background(), int32(numFeeds))
		if err != nil {
			log.Printf("Error fetching next feed batch: %v", err)
			continue
		}
		log.Printf("Found %v feeds. Fetching...", len(feedArr))

		wg := &sync.WaitGroup{}
		for _, v := range feedArr {
			wg.Add(1)
			go scrape(db, wg, v)
		}
		wg.Wait()
	}
}

func scrape(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed %v as fetched: %v", feed.ID, err)
		return
	}

	data, err := fetchRSS(feed.Url)
	if err != nil {
		log.Printf("Worker error: %s", err)
		panic(err)
	}
	for _, v := range data.Channel.Items {
		pubDate := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, v.PubDate); err == nil {
			pubDate = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     v.Title,
			Url:       v.Link,
			Description: sql.NullString{
				String: v.Description,
				Valid:  true,
			},
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		})

		if err != nil {
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
}
