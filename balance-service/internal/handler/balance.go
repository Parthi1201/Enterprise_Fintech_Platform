package handler

import "balance-service/internal/data"

type BalanceUsecase struct {
	repo *data.BalanceRepo
}

func NewBalanceUsecase(repo *data.BalanceRepo) *BalanceUsecase {
	return &BalanceUsecase{repo: repo}
}
