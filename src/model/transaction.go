package model

type Transaction struct {
	Invoice_id int     `json:"invoice_id"`
	Amount     float64 `json:"amount"`
	Reference  string  `json:"reference"`
}
