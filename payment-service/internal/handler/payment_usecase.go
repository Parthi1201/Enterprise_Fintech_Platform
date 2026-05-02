package handler

import (
	"payment-service/internal/client"
	"payment-service/internal/data"
)

type PaymentUsecase struct {
	repo          *data.PaymentRepo
	accountClient *account.Client
}

func NewPaymentUsecase(repo *data.PaymentRepo,accountClient *account.Client) *PaymentUsecase {
	return &PaymentUsecase{repo:repo,accountClient: accountClient}
}
