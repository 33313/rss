package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/myshkovsky/rss/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func deserializeUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
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

func deserializeFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func deserializeFeedArray(feeds []database.Feed) []Feed {
	arr := make([]Feed, len(feeds))
	for i, feed := range feeds {
		arr[i] = deserializeFeed(feed)
	}
	return arr
}

type Follow struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func deserializeFollow(follow database.FeedFollow) Follow {
    return Follow{
    	ID:        follow.ID,
    	FeedID:    follow.FeedID,
    	UserID:    follow.UserID,
    	CreatedAt: follow.CreatedAt,
    	UpdatedAt: follow.UpdatedAt,
    }
}
