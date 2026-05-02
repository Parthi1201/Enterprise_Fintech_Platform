package data

import (
	"context"
	"transaction-service/model"

	"gorm.io/gorm"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(d *Data) *TransactionRepo {
	return &TransactionRepo{db: d.DB}
}

func (r *TransactionRepo) Create(ctx context.Context, tx *model.Transaction) error {
	return r.db.WithContext(ctx).Create(tx).Error
}

func (r *TransactionRepo) GetByID(ctx context.Context,transactionID int64) (*model.Transaction, error) {
	var tx model.Transaction
	err := r.db.WithContext(ctx).Where("transaction_id=?", transactionID).First(&tx).Error
	if err != nil {
		return nil, err
	}
	return &tx, nil
}


func (r *TransactionRepo) ListByAccountID(ctx context.Context,accountID int64) ([]*model.Transaction, error) {
	var txs []*model.Transaction

	err := r.db.WithContext(ctx).
		Where("account_id = ?", accountID).Order("creation_time DESC").Find(&txs).Error

	if err != nil {
		return nil, err
	}
	return txs, nil
}
