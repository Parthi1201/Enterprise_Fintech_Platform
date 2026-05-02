package handler

import (
	"context"

	"balance-service/internal/data"
)

func (uc *BalanceUsecase) GetCurrentBalance(
	ctx context.Context,
	accountID string,
) (*data.Balance, error) {

	return uc.repo.GetLatestByAccount(ctx, accountID)
}
