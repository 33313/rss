package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func decodeParams[T any](w http.ResponseWriter, r *http.Request, data *T) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
