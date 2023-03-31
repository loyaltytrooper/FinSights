package models

type ML_Response struct {
	Data string `json:"data"`
}

type MarkedResponse struct {
	TxnId      string `json:"trans_id"`
	Prediction string `json:"pred"`
	Score      string `json:"score"`
}

//type Markers
