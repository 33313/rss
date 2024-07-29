package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Postgres driver
	"github.com/myshkovsky/rss/internal/database"
)

func main() {
	// Config
	godotenv.Load()
	address := fmt.Sprintf("%s:%s", os.Getenv("DOMAIN"), os.Getenv("PORT"))
	db, err := sql.Open("postgres", os.Getenv("DB_CONN"))
	if err != nil {
		panic(err)
	}
	queries := database.New(db)
	api := API{
		DB:         queries,
		FetchLimit: 10,
	}
	mux := http.NewServeMux()

	// Routes
	// General
	mux.HandleFunc("GET /v1/healthz", api.healthz)
	mux.HandleFunc("GET /v1/err", api.err)
	// Users
	mux.HandleFunc("GET /v1/users", api.auth(api.UsersGet))
	mux.HandleFunc("POST /v1/users", api.UsersPost)
	// Feeds
	mux.HandleFunc("GET /v1/feeds", api.FeedsGet)
	mux.HandleFunc("POST /v1/feeds", api.auth(api.FeedsPost))
	// Follows
	mux.HandleFunc("POST /v1/feed_follows", api.auth(api.FollowsPost))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", api.FollowsDelete)
	mux.HandleFunc("GET /v1/feed_follows", api.auth(api.FollowsGetFromUser))

	go StartScraper(queries, 10, time.Minute)

	fmt.Println("Running server on", address)
	srv := http.Server{
		Addr:    address,
		Handler: mux,
	}
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
