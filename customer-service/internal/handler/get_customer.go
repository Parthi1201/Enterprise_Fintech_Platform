package handler

import (
	"context"
	"strconv"

	pb "customer-service/api/customer"
	"customer-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

//get
func (uc *CustomerUseCase) GetCustomer(ctx context.Context,req *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {

	customerID, err := strconv.ParseInt(req.CustomerId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest("INVALID_CUSTOMER_ID","customer_id must be a valid numeric ID")
	}
	customer, err := uc.repo.GetCustomerByID(ctx, customerID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound("CUSTOMER_NOT_FOUND","customer not found")
		}

		return nil, errors.InternalServer("GET_CUSTOMER_FAILED",err.Error())
	}
	return &pb.GetCustomerResponse{Customer: mapper.MapCustomerToProto(customer)}, nil
}
