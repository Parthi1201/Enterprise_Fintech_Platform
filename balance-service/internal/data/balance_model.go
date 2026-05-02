package data

import "time"

type Balance struct {
	AccountID string `gorm:"size:64;index;not null"`
	BalanceDate time.Time `gorm:"type:date;not null"`
	BalanceTimestamp time.Time `gorm:"not null"`
	AvailableBalance string `gorm:"type:decimal(20,4);not null"`
	PendingBalance string `gorm:"type:decimal(20,4);not null"`
	Currency string `gorm:"size:3;not null"`
	ChangeAmount string `gorm:"type:decimal(20,4);not null"`
	ChangeReason string `gorm:"size:50;not null"`
}
