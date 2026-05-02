package data

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type BalanceRepo struct {
	db *gorm.DB
}

func NewBalanceRepo(data *Data) *BalanceRepo {
	return &BalanceRepo{db:data.DB}
}


func (r *BalanceRepo) Create(
	ctx context.Context,
	balance *Balance,
) error {
	return r.db.WithContext(ctx).Create(balance).Error
}


func (r *BalanceRepo) GetLatestByAccount(
	ctx context.Context,
	accountID string,
) (*Balance, error) {

	var balance Balance

	err := r.db.WithContext(ctx).
		Where("account_id = ?", accountID).
		Order("balance_timestamp DESC").
		First(&balance).Error

	if err != nil {
		return nil, err
	}

	return &balance, nil
}

func (r *BalanceRepo) ListByAccount(
	ctx context.Context,
	accountID string,
	fromDate *time.Time,
	toDate *time.Time,
) ([]*Balance, error) {

	var balances []*Balance

	q := r.db.WithContext(ctx).
		Where("account_id = ?", accountID)

	if fromDate != nil {
		q = q.Where("balance_date >= ?", *fromDate)
	}
	if toDate != nil {
		q = q.Where("balance_date <= ?", *toDate)
	}

	err := q.
		Order("balance_timestamp DESC").
		Find(&balances).Error

	if err != nil {
		return nil, err
	}

	return balances, nil
}
