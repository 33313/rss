package main

import (
	"log"
	"net/http"

	"github.com/33313/rss/internal/auth"
	"github.com/33313/rss/internal/database"
)

type AuthHandler func(http.ResponseWriter, *http.Request, database.User)

func (api *API) auth(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := auth.StripHeader(r)
		if err != nil {
			log.Printf("Error stripping header: %s", err)
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		user, err := api.DB.GetUserByApiKey(r.Context(), key)
		if err != nil {
			log.Printf("Error getting user: %s", err)
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		handler(w, r, user)
	}
}
