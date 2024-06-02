package models

type User struct {
	Balance   float64 `json:"balance"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	UserId    int     `json:"user_id"`
}
