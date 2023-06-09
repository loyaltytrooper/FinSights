package services

import (
	"FinSights/bank_pdf_parsers"
	"FinSights/helpers"
	"FinSights/models"
)

func GetFundTrail(fileName string, password string) models.TransactionsJSON {
	// TODO remove password dependency
	txns, csvData := bank_pdf_parsers.ParseAxisPDF(fileName, password)

	go helpers.CreateCSV(csvData, fileName[0:len(fileName)-4])

	jsonTxns := helpers.ChangeToJSON(&txns)
	jsonResult := models.TransactionsJSON{
		Id:   "-1",
		Name: "Transactions",
		Position: models.Position{
			X: 0,
			Y: 10000,
		},
		MetaData: models.NodeData{
			TransferMode: "custom",
		},
	}
	for _, v := range jsonTxns {
		jsonResult.MetaData.Children = append(jsonResult.MetaData.Children, v)
	}
	return jsonResult
}
