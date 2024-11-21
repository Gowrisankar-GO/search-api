package model

type Users struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	PhoneNumber string  `json:"phone_number"`
	Country     string  `json:"country"`
	Score       float64 `json:"score"`
}
