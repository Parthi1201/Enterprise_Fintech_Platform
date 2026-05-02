package handler

import (
	"context"
	"strings"
	"time"

	pb "customer-service/api/customer"
	"customer-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (uc *CustomerUseCase) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	if strings.TrimSpace(req.CustomerNumber) == "" {
		return nil, errors.BadRequest("CUSTOMER_NUMBER_REQUIRED", "customer_number is required")
	}

	if req.DateOfBirth != "" {
		if _, err := time.Parse("2006-01-02", req.DateOfBirth); err != nil {
			return nil, errors.BadRequest("INVALID_DATE_OF_BIRTH", "date_of_birth must be YYYY-MM-DD")
		}
	}

	customer := mapper.HandlerMapper(req)

	if err := uc.repo.Create(ctx, customer); err != nil {

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.Conflict("DUPLICATE_CUSTOMER_NUMBER", "this customer number already exists")
		}
		return nil, errors.BadRequest("CREATE_CUSTOMER_FAILED", err.Error())
	}

	return &pb.CreateCustomerResponse{Customer: mapper.MapCustomerToProto(customer)}, nil
}
