package helpers

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CreateCSV(csvData map[string][]string, fileName string) {
	// creating a csv file
	csvFile, err := os.Create(fileName + "transactions.csv")
	if err != nil {
		fmt.Println(err.Error() + "from ml_csv_creator")
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			fmt.Println(err.Error() + "from ML_CSV_Creator")
		}
	}(csvFile)

	// writing the headers
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()
	err = writer.Write([]string{"Date", "Type", "Category", "Destination", "Transaction ID", "Difference", "Previous Balance", "Final Balance"})
	if err != nil {
		fmt.Println(err.Error() + "from ML_CSV_Creator")
	}

	// writing the data
	for _, value := range csvData {
		err := writer.Write(value)
		if err != nil {
			fmt.Println(err.Error() + "from ML_CSV_Creator")
		}
	}
}
