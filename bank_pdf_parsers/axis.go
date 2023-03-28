package bank_pdf_parsers

import (
	helper "FinSights/helpers"
	"FinSights/models"
	"bufio"
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func ParseAxisPDF(fileName string, password string) (models.Transactions, map[string][]string) {
	os.Mkdir("act_"+fileName[0:len(fileName)-4], 0777)
	err := api.ExtractContentFile(fileName, "act_"+fileName[0:len(fileName)-4], nil, model.NewAESConfiguration(password, password, 128))
	if err != nil {
		fmt.Println(err.Error())
	}
	file, err := os.Open("act_" + fileName[0:len(fileName)-4])
	files, err := file.ReadDir(0)

	var transactions models.Transactions
	csvData := map[string][]string{}

	for _, f := range files {
		if !f.IsDir() {
			ReadAxisFile("act_"+fileName[0:len(fileName)-4]+"/"+f.Name(), &transactions, &csvData)
		}
	}
	return transactions, csvData
}

func ReadAxisFile(file string, transactions *models.Transactions, csvData *map[string][]string) (closingFound bool) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Error closing file")
		}
	}(f)

	// in case of words that ended with j and had length less than 5 or 6
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("Panic occurred")
		}
	}()

	//var txn models.Transaction
	var tempData []string
	foundTable := false
	var prevBalance float64

	// reading from the file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text()[len(scanner.Text())-1] == 'j' {
			word := scanner.Text()[1 : len(scanner.Text())-4]
			if strings.Compare(word, "Opening Balance") == 0 {
				for scanner.Scan() {
					if scanner.Text()[len(scanner.Text())-1] != 'j' {
						continue
					} else {
						tempBalance, err := strconv.ParseFloat(scanner.Text()[1:len(scanner.Text())-4], 64)
						if err == nil {
							prevBalance = tempBalance
						} else {
							panic(err)
						}
						break
					}
				}
				continue
			}

			if strings.Compare(word, "Closing Balance") == 0 && foundTable == true {
				foundTable = false
				txnTime, err := time.Parse("02-01-2006", tempData[0])
				if len(tempData) >= 4 && err == nil {
					difference, err := helper.SimplifyCommaNumber(tempData[len(tempData)-2])
					if err != nil {
						tempData = nil
					} else {
						balance, err := helper.SimplifyCommaNumber(tempData[len(tempData)-1])
						if err != nil {
							tempData = nil
						} else {
							descriptions := helper.GetDescription(tempData)
							id, mode, dst := helper.TxnDescription(descriptions)
							if balance > prevBalance {
								(*csvData)[id] = []string{txnTime.Format("02-01-2006"), "CREDIT", id, mode, dst, fmt.Sprintf("%v", difference), fmt.Sprintf("%v", prevBalance), fmt.Sprintf("%v", balance)}
								transactions.Txns = append(transactions.Txns, models.Transaction{
									Date:         txnTime,
									TxnType:      "CREDIT",
									TxnId:        id,
									TransferMode: mode,
									Destination:  dst,
									Difference:   difference,
									FinalAmount:  balance,
								})
							} else if balance < prevBalance {
								(*csvData)[id] = []string{txnTime.Format("02-01-2006"), "DEBIT", id, mode, dst, fmt.Sprintf("%v", difference), fmt.Sprintf("%v", prevBalance), fmt.Sprintf("%v", balance)}
								transactions.Txns = append(transactions.Txns, models.Transaction{
									Date:         txnTime,
									TxnType:      "DEBIT",
									TxnId:        id,
									TransferMode: mode,
									Destination:  dst,
									Difference:   -(difference),
									FinalAmount:  balance,
								})
							}
							prevBalance = balance
						}
					}
				}
				tempData = nil
			}

			txnTime, err := time.Parse("02-01-2006", word)
			if err != nil && foundTable == true {
				tempData = append(tempData, word)
				continue
			} else if err == nil {
				if len(tempData) >= 4 {
					difference, err := helper.SimplifyCommaNumber(tempData[len(tempData)-2])
					if err != nil {
						tempData = nil
					} else {
						balance, err := helper.SimplifyCommaNumber(tempData[len(tempData)-1])
						if err != nil {
							tempData = nil
						} else {
							descriptions := helper.GetDescription(tempData)
							id, mode, dst := helper.TxnDescription(descriptions)
							if balance > prevBalance {
								(*csvData)[id] = []string{txnTime.Format("02-01-2006"), "CREDIT", id, mode, dst, fmt.Sprintf("%v", difference), fmt.Sprintf("%v", prevBalance), fmt.Sprintf("%v", balance)}
								transactions.Txns = append(transactions.Txns, models.Transaction{
									Date:         txnTime,
									TxnId:        id,
									TransferMode: mode,
									Destination:  dst,
									Difference:   difference,
									FinalAmount:  balance,
								})
							} else if balance < prevBalance {
								(*csvData)[id] = []string{txnTime.Format("02-01-2006"), "DEBIT", id, mode, dst, fmt.Sprintf("%v", difference), fmt.Sprintf("%v", prevBalance), fmt.Sprintf("%v", balance)}
								transactions.Txns = append(transactions.Txns, models.Transaction{
									Date:         txnTime,
									TxnId:        id,
									TransferMode: mode,
									Destination:  dst,
									Difference:   -(difference),
									FinalAmount:  balance,
								})
							}
							prevBalance = balance
						}
					}
				}
				tempData = nil
				foundTable = true
				tempData = append(tempData, word)
				continue
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
