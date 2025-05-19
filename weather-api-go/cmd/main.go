package main

import (
	"log"
	"net/http"
	"weather-api-go/internal"
	"weather-api-go/internal/api"
)

func main() {
	internal.InitDB("weather.db")
	router := api.NewRouter()
	log.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
