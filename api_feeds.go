package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/myshkovsky/rss/internal/database"
)

func (api *API) FeedsPost(w http.ResponseWriter, r *http.Request, user database.User) {
	type Params struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := Params{}
	err := decodeParams[Params](w, r, &params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	feed, err := api.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:            uuid.New(),
		Name:          params.Name,
		Url:           params.Url,
		UserID:        user.ID,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
		LastFetchedAt: sql.NullTime{},
	})

	if err != nil {
		log.Printf("Error creating feed: %s", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	follow, err := api.DB.CreateFollow(r.Context(), database.CreateFollowParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    feed.UserID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		log.Printf("Error creating follow: %s", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, joinFeedFollow(feed, follow))
}

func (api *API) FeedsGet(w http.ResponseWriter, r *http.Request) {
	feeds, err := api.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, deserializeFeedArray(feeds))
}
