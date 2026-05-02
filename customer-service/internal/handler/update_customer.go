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

func (uc *CustomerUseCase) UpdateCustomer(ctx context.Context,req *pb.UpdateCustomerRequest) (*pb.UpdateCustomerResponse, error) {

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
	if req.FirstName != "" {
		customer.FirstName = req.FirstName
	}
	if req.LastName != "" {
		customer.LastName = req.LastName
	}
	if req.Phone != "" {
		customer.Phone = req.Phone
	}
	if req.Email != "" {
		customer.Email = req.Email
	}
	if req.DateOfBirth != "" {
		t, err := time.Parse("2006-01-02", req.DateOfBirth)
		if err != nil {
			return nil, errors.BadRequest("INVALID_DATE_OF_BIRTH","date_of_birth must be YYYY-MM-DD")
		}
		customer.DateOfBirth = t.String()
	}
	if err := uc.repo.Update(ctx, customer); err != nil {
		return nil, errors.InternalServer("UPDATE_CUSTOMER_FAILED",err.Error())
	}
	return &pb.UpdateCustomerResponse{Customer: mapper.MapCustomerToProto(customer)}, nil
}
