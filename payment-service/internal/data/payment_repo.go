package data

import (
	"context"

	"payment-service/internal/model"

	"gorm.io/gorm"
)

type PaymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

// Create
func (r *PaymentRepo) Create(ctx context.Context, p *model.Payment) error {
	return r.db.WithContext(ctx).Create(p).Error
}

// Get by ID
func (r *PaymentRepo) GetByID(ctx context.Context, paymentID int64) (*model.Payment, error) {
	var p model.Payment
	if err := r.db.WithContext(ctx).
		Where("payment_id = ?", paymentID).
		First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

// List by account (from OR to)
func (r *PaymentRepo) ListByAccountID(
	ctx context.Context,
	accountID int64,
) ([]*model.Payment, error) {

	var payments []*model.Payment
	if err := r.db.WithContext(ctx).
		Where("from_account_id = ? OR to_account_id = ?", accountID, accountID).
		Order("creation_time DESC").
		Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

// Update

func (r *PaymentRepo) Update(ctx context.Context, payment *model.Payment) error {
	return r.db.WithContext(ctx).Save(payment).Error
}
