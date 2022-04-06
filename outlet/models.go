package outlet

import (
	"jobapp.com/m/common"
	"jobapp.com/m/transaction"
)

type OutletModel struct {
	OutletName   string                         `json:"outlet_name" bson:"outlet_model" binding:"required"`
	Transactions []transaction.TransactionModel `json:"transactions" bson:"transactions" binding:"required"`
	Modified     common.ModifiedModel           `json:"modified" bson:"modified"`
}
