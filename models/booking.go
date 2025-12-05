package models

type Booking struct {
	ID           int    `json:"id" db:"id"`
	CourtID      int    `json:"court_id" db:"court_id"`
	UserID       int    `json:"user_id" db:"user_id"`
	CustomerName string `json:"customer_name" db:"customer_name"`
	BookingDate  string `json:"booking_date" db:"booking_date"`
	StartTime    string `json:"start_time" db:"start_time"`
	EndTime      string `json:"end_time" db:"end_time"`
	TotalPrice   int    `json:"total_price" db:"total_price"`
}