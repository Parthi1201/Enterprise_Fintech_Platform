package handler

import (
	"account-service/internal/client"
	"account-service/internal/data"
)

type AccountUsecase struct {
	repo *data.AccountRepo
	customerClient *client.CustomerClient
}

func NewAccountUsecase(repo *data.AccountRepo,customerClient *client.CustomerClient) *AccountUsecase {
	return &AccountUsecase{repo: repo,customerClient: customerClient}
}
