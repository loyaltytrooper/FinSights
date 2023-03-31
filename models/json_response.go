package models

type TransactionJSON struct {
	Date         string            `json:"txn_time"`
	TxnType      string            `json:"txn_type"`
	TxnId        string            `json:"id"`
	TransferMode string            `json:"payment_mode"`
	Suspicious   bool              `json:"suspicious"`
	ParentId     string            `json:"parent"`
	Destination  string            `json:"destination"`
	Difference   float64           `json:"name"`
	FinalAmount  float64           `json:"final_amount"`
	Children     []TransactionJSON `json:"children"`
}

type TransactionsJSON struct {
	Id   string            `json:"id"`
	Name string            `json:"name"`
	Txns []TransactionJSON `json:"children"`
}
