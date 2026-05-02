package mapper

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	snowid "Enterprise_Fintech_Platform/common/id"

	"card-service/internal/model"
)

func NewCardModel(accountID int64,cardNumber string,cardType string) *model.Card {

	hash := sha256.Sum256([]byte(cardNumber))
	cardNumberHash := hex.EncodeToString(hash[:])

	now := time.Now()
	id := snowid.New()

	return &model.Card{
		CardID:         int64(id),
		AccountID:      accountID,
		CardNumberHash: cardNumberHash,
		CardType:       cardType,
		CardStatus:     "active",
		ExpiryDate:     now.AddDate(1, 0, 0),
		DailyLimit:     0,
		MonthlyLimit:   0,
		CreationTime:   now,
		ModifiedTime:   now,
	}
}
