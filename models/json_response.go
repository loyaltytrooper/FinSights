package models

type TransactionJSON struct {
	TxnId    string   `json:"id"`
	ParentId string   `json:"parent"`
	MetaData NodeData `json:"data"`
}

type NodeData struct {
	Destination  string            `json:"destination"`
	Difference   float64           `json:"name"`
	FinalAmount  float64           `json:"final_amount"`
	TransferMode string            `json:"payment_mode"`
	Suspicious   bool              `json:"suspicious"`
	Date         string            `json:"txn_time"`
	TxnType      string            `json:"txn_type"`
	Children     []TransactionJSON `json:"children"`
}

type TransactionsJSON struct {
	Id       string   `json:"id"`
	Name     string   `json:"type"`
	Position Position `json:"position"`
	MetaData NodeData `json:"data"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}
