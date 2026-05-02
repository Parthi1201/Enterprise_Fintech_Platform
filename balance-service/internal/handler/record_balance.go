package handler

import (
	"context"
	"time"

	"balance-service/internal/data"
)

func (uc *BalanceUsecase) RecordBalanceChange(
	ctx context.Context,
	accountID string,
	available string,
	pending string,
	currency string,
	changeAmount string,
	changeReason string,
) (*data.Balance, error) {

	now := time.Now()

	balance := &data.Balance{
		AccountID:        accountID,
		AvailableBalance: available,
		PendingBalance:   pending,
		Currency:         currency,

		ChangeAmount: changeAmount,
		ChangeReason: changeReason,

		BalanceTimestamp: now,
		BalanceDate:      now,
	}

	if err := uc.repo.Create(ctx, balance); err != nil {
		return nil, err
	}

	return balance, nil
}
