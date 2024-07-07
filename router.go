package main

import "net/http"

type Router struct {}

func (router *Router) healthz(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct {
		Status string `json:"status"`
	}{Status: "ok"})
}

func (router *Router) err(w http.ResponseWriter, r *http.Request) {
    respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
