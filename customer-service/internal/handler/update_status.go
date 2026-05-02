package handler

import (
	"context"
	"strconv"

	pb "customer-service/api/customer"
	"customer-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (uc *CustomerUseCase) UpdateCustomerStatus(ctx context.Context,req *pb.UpdateCustomerStatusRequest) (*pb.Customer, error) {

	customerID, err := strconv.ParseInt(req.CustomerId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest("INVALID_CUSTOMER_ID","customer_id must be a valid numeric ID")
	}

	switch req.Status {
	case "active", "suspended", "closed":
	default:
		return nil, errors.BadRequest("INVALID_CUSTOMER_STATUS","status must be active, suspended, or closed")
	}

	customer, err := uc.repo.GetCustomerByID(ctx, customerID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound("CUSTOMER_NOT_FOUND","customer not found")
		}
		return nil, errors.InternalServer("GET_CUSTOMER_FAILED",err.Error())
	}

	if customer.Status == "closed" {
		return nil, errors.Conflict("CUSTOMER_ALREADY_CLOSED","closed customer cannot be modified")
	}

	customer.Status = req.Status

	if err := uc.repo.Update(ctx, customer); err != nil {
		return nil, errors.InternalServer("UPDATE_CUSTOMER_STATUS_FAILED",err.Error())
	}

	return mapper.MapCustomerToProto(customer), nil
}
