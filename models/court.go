package models

type Court struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	PricePerHour int    `json:"price_per_hour"`
	IsAvailable  bool   `json:"is_available"`
}