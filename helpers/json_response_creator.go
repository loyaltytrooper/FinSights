package helpers

import (
	"FinSights/models"
	"sort"
	"strconv"
)

func ChangeToJSON(txns *models.Transactions) map[string]models.TransactionJSON {
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
	finalTxns[accTxn.Destination] = accTxn
	return finalTxns
}

func AddChildrenRecursively(currentChild *models.TransactionJSON, temp_txns []models.Transaction, i int, parentID string) {
	if (len(temp_txns)) == i {
		return
	} else {
		(*currentChild).Children = append((*currentChild).Children, models.TransactionJSON{
			Date:         (temp_txns)[i].Date,
			TxnType:      (temp_txns)[i].TxnType,
			TxnId:        (temp_txns)[i].TxnId,
			ParentId:     parentID,
			TransferMode: (temp_txns)[i].TransferMode,
			Destination:  (temp_txns)[i].Destination,
			Difference:   (temp_txns)[i].Difference,
			FinalAmount:  (temp_txns)[i].FinalAmount,
		})
		AddChildrenRecursively(&currentChild.Children[len((*currentChild).Children)-1], temp_txns, i+1, (temp_txns)[i].TxnId)
	}
}
