package mapper

import (
	"account-service/internal/data"
	"account-service/model"

	"github.com/shopspring/decimal"
)

func ToDomainAccount(m *model.Account)*data.Account {
	if m==nil{
		return nil
	}

	var creditLimit *decimal.Decimal
	if m.CreditLimit!=0 {
		cl:=decimal.NewFromFloat(m.CreditLimit)
		creditLimit = &cl
	}

	return &data.Account{
		AccountID:m.AccountID,
		CustomerID:m.CustomerID,
		AccountNumber:m.AccountNumber,
		AccountType:m.AccountType,
		Currency:m.Currency,
		Status:m.Status,
		AvailableBalance:decimal.NewFromFloat(m.AvailableBalance),
		PendingBalance:decimal.NewFromFloat(m.PendingBalance),
		CreditLimit:creditLimit,
		LastTransactionAt:&m.LastTransactionAt,
		CreationTime:&m.CreationTime,
		ModifiedTime:&m.ModifiedTime,
	}
}
