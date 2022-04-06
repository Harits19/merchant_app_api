package merchant

import (
	"jobapp.com/m/common"
	"jobapp.com/m/outlet"
	"jobapp.com/m/transaction"
)

type MerchantModel struct {
	UserId       int64                          `json:"user_id" bson:"user_id" binding:"required"`
	MerchantName string                         `json:"merchant_name" bson:"merchant_name" binding:"required"`
	Outlets      []outlet.OutletModel           `json:"outlets" bson:"outlets" binding:"required"`
	Transactions []transaction.TransactionModel `json:"transactions" bson:"transactions" binding:"required"`
	Modified     common.ModifiedModel           `json:"modified" bson:"modified"`
}
