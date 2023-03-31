package services

import (
	"FinSights/helpers"
	"FinSights/models"
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func ParseCSVFile(fileName string) models.TransactionsJSON {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("(")
		fmt.Println(err.Error() + "from ParseCSV")
	}

	fileReader := csv.NewReader(bufio.NewReader(file))
	fileReader.Comma = ';'
	fileReader.TrimLeadingSpace = true

	records, err2 := fileReader.ReadAll()

	if err2 != nil {
		fmt.Println("@")
		fmt.Println(err2.Error() + "from ParseCSV")
	}

	var txns models.TransactionsPDF

	for _, record := range records {
		difference, err := helpers.SimplifyCommaNumber(record[5])
		if err != nil {
			fmt.Println("#")
			fmt.Println(err.Error() + "from ParseCSV")
		}
		finalAmount, err := helpers.SimplifyCommaNumber(record[6])
		if err != nil {
			fmt.Println("*")
			fmt.Println(err.Error() + "from ParseCSV")
		}

		txns.Txns = append(txns.Txns, models.Transaction{
			Date:         record[2],
			TxnType:      record[3],
			TxnId:        record[0],
			TransferMode: record[4],
			Destination:  record[1],
			Difference:   difference,
			FinalAmount:  finalAmount,
		})
	}

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
