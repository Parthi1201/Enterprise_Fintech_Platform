package handler

import (
	"context"
	"strconv"
	"time"

	pb "customer-service/api/customer"
	"customer-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (uc *CustomerUseCase) DiscardCustomer(ctx context.Context, req *pb.DiscardCustomerRequest) (*pb.Customer, error) {

	customerID, err := strconv.ParseInt(req.CustomerId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest("INVALID_CUSTOMER_ID", "customer_id must be a valid numeric ID")
	}
	customer, err := uc.repo.GetCustomerByID(ctx, customerID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound("CUSTOMER_NOT_FOUND", "customer not found")
		}
		return nil, errors.InternalServer("GET_CUSTOMER_FAILED", err.Error())
	}
	if customer.Discarded == "Y" {
		return nil, errors.Conflict("CUSTOMER_ALREADY_DISCARDED", "customer already discarded")
	}
	now := time.Now()
	customer.Discarded = "Y"
	customer.DiscardTime = now

	if err := uc.repo.Update(ctx, customer); err != nil {
		return nil, errors.InternalServer("DISCARD_CUSTOMER_FAILED",err.Error())
	}
	return mapper.MapCustomerToProto(customer), nil
}
