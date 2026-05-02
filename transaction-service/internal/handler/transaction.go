package handler

import (
	"transaction-service/internal/client"
	"transaction-service/internal/data"
)

type TransactionUsecase struct {
	repo *data.TransactionRepo
	accountClient *client.AccountClient

}

func NewTransactionUsecase(repo *data.TransactionRepo,accountClient *client.AccountClient) *TransactionUsecase {
	return &TransactionUsecase{repo:repo,accountClient: accountClient}
}

