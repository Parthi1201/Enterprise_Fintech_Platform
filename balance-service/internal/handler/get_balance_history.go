package handler

import (
	"context"
	"time"

	"balance-service/internal/data"
)

func (uc *BalanceUsecase) GetBalanceHistory(
	ctx context.Context,
	accountID string,
	fromDate *time.Time,
	toDate *time.Time,
) ([]*data.Balance, error) {

	return uc.repo.ListByAccount(ctx, accountID, fromDate, toDate)
}
