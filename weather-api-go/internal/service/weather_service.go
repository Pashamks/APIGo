package service

import (
	"errors"
	"sync"
)

type Weather struct {
	City    string  `json:"city"`
	TempC   float64 `json:"temp_c"`
	Weather string  `json:"weather"`
}

type Subscription struct {
	City    string
	Email   string
	Active  bool
}

type WeatherService struct {
	mu            sync.Mutex
	subscriptions []Subscription
}

func NewWeatherService() *WeatherService {
	return &WeatherService{
		subscriptions: []Subscription{},
	}
}

func (ws *WeatherService) GetWeather(city string) (Weather, error) {
	// Simulate fetching weather data
	if city == "" {
		return Weather{}, errors.New("city cannot be empty")
	}

	return Weather{
		City:    city,
		TempC:   22.5,
		Weather: "Sunny",
	}, nil
}

func (ws *WeatherService) Subscribe(city, email string) error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	for _, sub := range ws.subscriptions {
		if sub.City == city && sub.Email == email {
			return errors.New("already subscribed")
		}
	}

	ws.subscriptions = append(ws.subscriptions, Subscription{
		City:   city,
		Email:  email,
		Active: true,
	})

	return nil
}

func (ws *WeatherService) Unsubscribe(city, email string) error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	for i, sub := range ws.subscriptions {
		if sub.City == city && sub.Email == email {
			ws.subscriptions = append(ws.subscriptions[:i], ws.subscriptions[i+1:]...)
			return nil
		}
	}

	return errors.New("subscription not found")
}