package data

import (
	"card-service/internal/model"
	"context"

	"gorm.io/gorm"
)

type CardRepo struct{
	db *gorm.DB
}

func NewCardRepo(data *Data) *CardRepo{
	return &CardRepo{db:data.DB}
}

func(r *CardRepo)Create(ctx context.Context, card *model.Card)error{
	return r.db.WithContext(ctx).Create(card).Error
}
func (r *CardRepo) GetByID(ctx context.Context, cardID int64) (*model.Card,error) {
	var card model.Card
	if err:=r.db.WithContext(ctx).Where("card_id=?",cardID).First(&card).Error;err!=nil{
		return nil,err
	}
	return &card,nil
}


func (r *CardRepo) ListByAccountID(ctx context.Context,accountID int64) ([]*model.Card, error){
	var cards []*model.Card
	if err:=r.db.WithContext(ctx).Where("account_id=?",accountID).Order("creation_time DESC").Find(&cards).Error;err!=nil{
		return nil, err
	}

	return cards, nil
}

func (r *CardRepo) UpdateStatus(ctx context.Context,cardID int64,status string) error {

	return r.db.WithContext(ctx).Model(&model.Card{}).Where("card_id=?",cardID).Update("card_status",status).Error
}

func (r *CardRepo) UpdateLimit(ctx context.Context,cardID int64,dailyLimit float64,monthlyLimit float64) error {
	return r.db.WithContext(ctx).Model(&model.Card{}).Where("card_id=?",cardID).
		Updates(map[string]interface{}{
			"daily_limit":dailyLimit,
			"monthly_limit":monthlyLimit,
		}).Error
}
