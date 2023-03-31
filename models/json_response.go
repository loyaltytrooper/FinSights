package models

type TransactionJSON struct {
	TxnId    string            `json:"id"`
	ParentId string            `json:"parent"`
	Children []TransactionJSON `json:"children"`
	MetaData NodeData          `json:"data"`
}

type NodeData struct {
	Destination  string  `json:"destination"`
	Difference   float64 `json:"name"`
	FinalAmount  float64 `json:"final_amount"`
	TransferMode string  `json:"payment_mode"`
	Suspicious   bool    `json:"suspicious"`
	Date         string  `json:"txn_time"`
	TxnType      string  `json:"txn_type"`
}

type TransactionsJSON struct {
	Id       string            `json:"id"`
	Name     string            `json:"name"`
	Txns     []TransactionJSON `json:"children"`
	MetaData NodeData          `json:"data"`
}
