package data

import (
	"time"

	"github.com/shopspring/decimal"
)

type Account struct {
	AccountID int64
	CustomerID int64
	AccountNumber string
	AccountType string
	Currency string
	Status string

	AvailableBalance decimal.Decimal
	PendingBalance decimal.Decimal
	CreditLimit *decimal.Decimal

	LastTransactionAt *time.Time
	CreationTime *time.Time
	ModifiedTime *time.Time
	CreatedBy *string
	ModifiedBy *string
}
