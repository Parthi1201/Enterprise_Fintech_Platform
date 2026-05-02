package data

import (
	"account-service/model"
	"context"

	"gorm.io/gorm"
)

type AccountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(data *Data) *AccountRepo {
	return &AccountRepo{db: data.DB}
}

func (r *AccountRepo) Create(ctx context.Context, account *model.Account) error {
	return r.db.WithContext(ctx).Create(account).Error
}

func (r *AccountRepo) GetByID(ctx context.Context, id int64) (*model.Account, error) {
	var account model.Account
	if err:=r.db.WithContext(ctx).Where("account_id=?", id).First(&account).Error;err!=nil{
		return nil,err
	}
	return &account, nil
}


func (r *AccountRepo)ListByCustomerID(ctx context.Context,customerID int64)([]*model.Account, error) {

	var accounts []*model.Account
	if err:=r.db.WithContext(ctx).Where("customer_id=?",customerID).Find(&accounts).Error;err!=nil{
		return nil,err
	}

	return accounts,nil
}


func (r *AccountRepo) UpdateStatus(ctx context.Context,accountID int64,status string)error{

	return r.db.WithContext(ctx).Model(&model.Account{}).Where("account_id=?",accountID).Update("status",status).Error
}
