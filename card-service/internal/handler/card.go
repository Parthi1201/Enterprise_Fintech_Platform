package handler

import (
	"card-service/internal/client"
	"card-service/internal/data"
)

type CardUsecase struct {
	repo *data.CardRepo
	accountClient *client.AccountClient
}

func NewCardUsecase(repo *data.CardRepo,accountClient *client.AccountClient) *CardUsecase {
	return &CardUsecase{repo: repo,accountClient: accountClient}
}
