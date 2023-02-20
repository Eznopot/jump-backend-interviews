package model

type Invoice struct {
	User_id int     `json:"user_id"`
	Amount  float64 `json:"amount"`
	Label   string  `json:"label"`
}
