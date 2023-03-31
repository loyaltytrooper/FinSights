package models

type ML_Response struct {
	Data []MarkedResponse `json:"data"`
}

type MarkedResponse struct {
	TxnId      int     `json:"trans_id"`
	Prediction int     `json:"pred"`
	Score      float64 `json:"score"`
}

//type Markers
