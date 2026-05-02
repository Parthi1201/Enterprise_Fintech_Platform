package mapper

import (
	snowid "Enterprise_Fintech_Platform/common/id"
	"account-service/model"
)

func NewAccountModel(customerID int64,accountNumber string,accountType string,currency string,) *model.Account {

	id := snowid.New()

	return &model.Account{
		AccountID:        int64(id),
		CustomerID:       customerID,
		AccountNumber:    accountNumber,
		AccountType:      accountType,
		Currency:         currency,
		Status:           "active",
		AvailableBalance: 0,
		PendingBalance:   0,
	}
}
