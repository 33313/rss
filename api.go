package main

import (
	"net/http"

	"github.com/33313/rss/internal/database"
)

type API struct {
	DB         *database.Queries
	FetchLimit int32
}

func (api *API) healthz(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct {
		Status string `json:"status"`
	}{Status: "ok"})
}

func (api *API) err(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
