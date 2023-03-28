package models

import "time"

type Transaction struct {
	Date         time.Time `json:"txn_time"`
	TxnType      string    `json:"txn_type"`
	TxnId        string    `json:"id"`
	TransferMode string    `json:"payment_mode"`
	Suspicious   bool      `json:"suspicious"`
	Destination  string    `json:"destination"`
	Difference   float64   `json:"name"`
	FinalAmount  float64   `json:"final_amount"`
}

type Transactions struct {
	Txns []Transaction `json:"transactions"`
}
