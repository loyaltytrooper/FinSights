package helpers

import (
	"FinSights/models"
	"sort"
	"strconv"
)

func ChangeToJSON(txns *models.TransactionsPDF) map[string]models.TransactionJSON {
	finalTxns := make(map[string]models.TransactionJSON)
	sort.Slice(txns.Txns, func(i, j int) bool {
		return txns.Txns[i].Destination < txns.Txns[j].Destination
	})

	var temp_arr []models.Transaction
	for _, txn := range txns.Txns {
		if len(temp_arr) == 0 || txn.Destination == temp_arr[0].Destination {
			temp_arr = append(temp_arr, txn)
		} else {
			accTxn := models.TransactionJSON{
				Date:         temp_arr[0].Date,
				TxnType:      temp_arr[0].TxnType,
				TxnId:        temp_arr[0].TxnId,
				ParentId:     strconv.Itoa(-1),
				TransferMode: temp_arr[0].TransferMode,
				Destination:  temp_arr[0].Destination,
				Difference:   temp_arr[0].Difference,
				FinalAmount:  temp_arr[0].FinalAmount,
			}
			if len(temp_arr) > 1 {
				AddChildrenRecursively(&accTxn, temp_arr[1:], 0, temp_arr[0].TxnId)
			}
			temp_arr = nil
			temp_arr = append(temp_arr, txn)
			finalTxns[accTxn.Destination] = accTxn
		}
	}
	// TODO this has to be done only in case of PDFs
	//accTxn := models.TransactionJSON{
	//	Date:         temp_arr[0].Date,
	//	TxnType:      temp_arr[0].TxnType,
	//	TxnId:        temp_arr[0].TxnId,
	//	ParentId:     strconv.Itoa(-1),
	//	TransferMode: temp_arr[0].TransferMode,
	//	Destination:  temp_arr[0].Destination,
	//	Difference:   temp_arr[0].Difference,
	//	FinalAmount:  temp_arr[0].FinalAmount,
	//}
	//if len(temp_arr) > 1 {
	//	AddChildrenRecursively(&accTxn, temp_arr[1:], 0, temp_arr[0].TxnId)
	//}
	//temp_arr = nil
	//finalTxns[accTxn.Destination] = accTxn
	return finalTxns
}
