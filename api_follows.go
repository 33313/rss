package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/myshkovsky/rss/internal/database"
)

func (api *API) FollowsPost(w http.ResponseWriter, r *http.Request, user database.User) {
	type Params struct {
		FeedId  uuid.UUID `json:"feed_id"`
	}
	params := Params{}
	err := decodeParams[Params](w, r, &params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	follow, err := api.DB.CreateFollow(r.Context(), database.CreateFollowParams{
		ID:     uuid.New(),
		FeedID: params.FeedId,
		UserID: user.ID,
	})

	if err != nil {
		log.Printf("Error following feed: %s", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, deserializeFollow(follow))
}
