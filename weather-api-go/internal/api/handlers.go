package api

import (
	"encoding/json"
	"net/http"
	"sync"
	"weather-api-go/internal/service"
)

type WeatherResponse struct {
	City    string  `json:"city"`
	TempC   float64 `json:"temp_c"`
	Weather string  `json:"weather"`
}

type Subscription struct {
	Email string `json:"email"`
}

var subscriptions = struct {
	sync.RWMutex
	emails []string
}{}

var weatherService = service.NewWeatherService()

func GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "city parameter is required", http.StatusBadRequest)
		return
	}

	weather, err := weatherService.GetWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := WeatherResponse{
		City:    weather.City,
		TempC:   weather.TempC,
		Weather: weather.Weather,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	var sub Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	subscriptions.Lock()
	subscriptions.emails = append(subscriptions.emails, sub.Email)
	subscriptions.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sub)
}

func ConfirmSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for confirming subscription
}

func UnsubscribeHandler(w http.ResponseWriter, r *http.Request) {
	var sub Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	subscriptions.Lock()
	defer subscriptions.Unlock()

	for i, email := range subscriptions.emails {
		if email == sub.Email {
			subscriptions.emails = append(subscriptions.emails[:i], subscriptions.emails[i+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
