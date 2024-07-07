package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    res, err := json.Marshal(payload)
    if err != nil {
        log.Printf("Error while marshaling JSON: %s", err)
        w.WriteHeader(http.StatusInternalServerError)
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(res)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
    respondWithJSON(w, code, struct{
        Error string `json:"error"`
    }{
        Error: msg,
    })
}
