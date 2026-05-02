package data

import (
	"context"
	"customer-service/internal/model"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type CustomerRepo struct{
	db  *gorm.DB
	log *log.Helper
}

func NewCustomerRepo(data *Data,logger log.Logger)*CustomerRepo {
	return &CustomerRepo{db:data.DB,log:log.NewHelper(logger)}
}
//create
func(r *CustomerRepo)Create(ctx context.Context,c *model.Customer)error{
	return r.db.WithContext(ctx).Create(c).Error
}

//get
func(r *CustomerRepo)GetCustomerByID(ctx context.Context,customerID int64)(*model.Customer, error) {
	var customer model.Customer
	err := r.db.WithContext(ctx).Where("customer_id = ?", customerID).First(&customer).Error
	if err != nil{
		return nil,err
	}
	return &customer, nil
}


//update
func(r *CustomerRepo)Update(ctx context.Context,customer *model.Customer)error{
	return r.db.WithContext(ctx).Save(customer).Error
}


