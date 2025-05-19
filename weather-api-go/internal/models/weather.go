package models

type Weather struct {
	City    string  `json:"city"`
	TempC   float64 `json:"temp_c"`
	Weather string  `json:"weather"`
}

type Subscription struct {
	Email    string `json:"email"`
	City     string `json:"city"`
	Active   bool   `json:"active"`
}