package services

import (
	"FinSights/bank_pdf_parsers"
	"FinSights/helpers"
	"FinSights/models"
)

func GetFundTrail(fileName string, password string) models.TransactionsJSON {
	// TODO remove password dependency
	txns, csvData := bank_pdf_parsers.ParseAxisPDF(fileName, password)

	go helpers.CreateCSV(csvData)

	jsonTxns := helpers.ChangeToJSON(&txns)
	jsonResult := models.TransactionsJSON{
		Id:   "-1",
		Name: "Transactions",
	}
	for _, v := range jsonTxns {
		jsonResult.Txns = append(jsonResult.Txns, v)
	}
	return jsonResult
}
