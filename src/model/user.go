package model

type User struct {
	Id         int     `json:"user_id"`
	First_name string  `json:"first_name"`
	Last_name  string  `json:"last_name"`
	Balance    float64 `json:"balance"`
}
