package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/weather", GetWeatherHandler).Methods("GET")
	router.HandleFunc("/subscribe", SubscribeHandler).Methods("POST")
	router.HandleFunc("/confirm-subscription", ConfirmSubscriptionHandler).Methods("POST")
	router.HandleFunc("/unsubscribe", UnsubscribeHandler).Methods("DELETE")

	return router
}