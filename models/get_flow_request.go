package models

type GetFlowRequest struct {
	Password    string  `json:"password"`
	FileUrl     string  `json:"file"`
	CreditScore float32 `json:"credit_score"`
}
