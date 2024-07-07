package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    godotenv.Load()
    address := fmt.Sprintf("%s:%s", os.Getenv("DOMAIN"), os.Getenv("PORT"))
    mux := http.NewServeMux()
	srv := http.Server{
        Addr: address,
		Handler: mux,
	}

	fmt.Println("Running server on", address)
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
