package main

import "github.com/myshkovsky/rss/internal/database"

type FeedFollowJoined struct {
	FeedObject   Feed   `json:"feed"`
	FollowObject Follow `json:"feed_follow"`
}

func joinFeedFollow(feed database.Feed, follow database.FeedFollow) FeedFollowJoined {
	return FeedFollowJoined{
		FeedObject:   deserializeFeed(feed),
		FollowObject: deserializeFollow(follow),
	}
}
