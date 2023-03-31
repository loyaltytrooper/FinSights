package models

type Transaction struct {
	Date         string  `json:"txn_time"`
	TxnType      string  `json:"txn_type"`
	TxnId        string  `json:"id"`
	TransferMode string  `json:"payment_mode"`
	Suspicious   bool    `json:"suspicious"`
	Destination  string  `json:"destination"`
	Difference   float64 `json:"name"`
	FinalAmount  float64 `json:"final_amount"`
}

type TransactionsPDF struct {
	Txns []Transaction `json:"transactions"`
}
