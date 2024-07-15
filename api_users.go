package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/myshkovsky/rss/internal/database"
)

func (api *API) UsersPost(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Name string `json:"name"`
	}
	params := Params{}
    err := decodeParams[Params](w, r, &params)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

	user, err := api.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJSON(w, http.StatusCreated, deserializeUser(user))
}

func (api *API) UsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
    respondWithJSON(w, http.StatusOK, deserializeUser(user))
}
