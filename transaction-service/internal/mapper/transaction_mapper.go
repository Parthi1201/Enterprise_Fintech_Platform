package mapper

import (
	"strings"
	"time"

	"transaction-service/model"
)

func NewTransactionModel(
	id int64,
	accountID int64,
	txType string,
	amount float64,
	currency string,
	description string,
	now time.Time,
) *model.Transaction {

	return &model.Transaction{
		TransactionID:        id,
		AccountID:            accountID,
		RelatedTransactionID: 0,                  // 0 → no related tx
		TransactionType:      strings.ToUpper(txType),
		Amount:               amount,
		Currency:             currency,
		Status:               "SUCCESS",          // ONLY SUCCESS / FAILED
		Description:          description,
		ReferenceNumber:      "",
		MerchantInfo:         "{}",               // MUST be valid JSON
		TransactionTime:      now,
		CreatedAt:            now,
		PostedAt:             time.Time{},        // zero value (acts as NULL)
		ArchivedAt:           time.Time{},
	}
}
