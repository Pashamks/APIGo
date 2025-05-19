package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type WeatherServiceInterface interface {
	GetWeather(city string) (WeatherResponse, error)
}

type mockWeatherService struct{}

func (m *mockWeatherService) GetWeather(city string) (WeatherResponse, error) {
	return WeatherResponse{City: city, TempC: 20.0, Weather: "Sunny"}, nil
}

var _ WeatherServiceInterface = (*mockWeatherService)(nil)

func TestGetWeatherHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/weather?city=Kyiv", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	GetWeatherHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resp WeatherResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("error decoding response: %v", err)
	}
	if resp.City != "Kyiv" {
		t.Errorf("expected city Kyiv, got %s", resp.City)
	}
}

func TestGetWeatherHandler_BadRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "/weather", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetWeatherHandler)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rr.Code)
	}
}

func TestSubscribeHandler(t *testing.T) {
	body := `{"email":"test@example.com"}`
	req, _ := http.NewRequest("POST", "/subscribe", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SubscribeHandler)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", rr.Code)
	}
}
