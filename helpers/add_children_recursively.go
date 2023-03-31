package helpers

import "FinSights/models"

func AddChildrenRecursively(currentChild *models.TransactionJSON, temp_txns []models.Transaction, i int, parentID string) {
	if (len(temp_txns)) == i {
		return
	} else {
		(*currentChild).MetaData.Children = append((*currentChild).MetaData.Children, models.TransactionJSON{
			TxnId:    (temp_txns)[i].TxnId,
			ParentId: parentID,
			MetaData: models.NodeData{
				Date:         (temp_txns)[i].Date,
				TxnType:      (temp_txns)[i].TxnType,
				TransferMode: (temp_txns)[i].TransferMode,
				Destination:  (temp_txns)[i].Destination,
				Difference:   (temp_txns)[i].Difference,
				FinalAmount:  (temp_txns)[i].FinalAmount,
			},
		})
		AddChildrenRecursively(&currentChild.MetaData.Children[len((*currentChild).MetaData.Children)-1], temp_txns, i+1, (temp_txns)[i].TxnId)
	}
}
