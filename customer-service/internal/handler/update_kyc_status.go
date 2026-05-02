package handler

import (
	"context"
	"strconv"

	pb "customer-service/api/customer"
	"customer-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (uc *CustomerUseCase) UpdateKycStatus(ctx context.Context,req *pb.UpdateKycStatusRequest) (*pb.Customer, error) {


	customerID, err := strconv.ParseInt(req.CustomerId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest("INVALID_CUSTOMER_ID","customer_id must be a valid numeric ID")
	}


	switch req.KycStatus {
	case "pending", "verified", "rejected":
	default:
		return nil, errors.BadRequest("INVALID_KYC_STATUS","kyc_status must be pending, verified, or rejected")
	}

	customer, err := uc.repo.GetCustomerByID(ctx, customerID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound("CUSTOMER_NOT_FOUND","customer not found")
		}
		return nil, errors.InternalServer("GET_CUSTOMER_FAILED",err.Error())
	}

	customer.KycStatus = req.KycStatus

	if err := uc.repo.Update(ctx, customer); err != nil {
		return nil, errors.InternalServer("UPDATE_KYC_STATUS_FAILED",err.Error())
	}

	return mapper.MapCustomerToProto(customer), nil
}
