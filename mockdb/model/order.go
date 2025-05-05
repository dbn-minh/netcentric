package model

type Order struct {
	ID     int         `json:"id"`
	UserID int         `json:"user_id"`
	Items  []OrderItem `json:"items"`
	Total  float64     `json:"total"`
}
