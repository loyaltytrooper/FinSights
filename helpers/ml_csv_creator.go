package helpers

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CreateCSV(csvData map[string][]string) {
	// creating a csv file
	csvFile, err := os.Create("transactions.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	// writing the headers
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()
	writer.Write([]string{"Date", "Txn Type", "Txn Id", "Transfer Mode", "Destination", "Difference", "Previous Balance", "Final Balance"})

	// writing the data
	for _, value := range csvData {
		err := writer.Write(value)
		if err != nil {
			fmt.Println(err)
		}
	}
}
