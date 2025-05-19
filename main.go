package main

import (
	"encoding/json"
	"net/http"
)

type WeatherResponse struct {
	City    string  `json:"city"`
	TempC   float64 `json:"temp_c"`
	Weather string  `json:"weather"`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "city parameter is required", http.StatusBadRequest)
		return
	}

	resp := WeatherResponse{
		City:    city,
		TempC:   22.5,
		Weather: "Sunny",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/weather", weatherHandler)
	http.ListenAndServe(":8080", nil)
}
