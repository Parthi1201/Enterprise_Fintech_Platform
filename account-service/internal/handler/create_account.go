package handler

import (
	"context"
	"strconv"
	"strings"

	pb "account-service/api/account"
	"account-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (uc *AccountUsecase) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {

	customerID, err := strconv.ParseInt(req.CustomerId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest("INVALID_CUSTOMER_ID", "customer_id must be a valid numeric ID")
	}

	if err := uc.customerClient.GetCustomer(
		ctx,
		strconv.FormatInt(customerID, 10),
	); err != nil {
		return nil, errors.NotFound("CUSTOMER_NOT_FOUND","customer does not exist")
	}

	if strings.TrimSpace(req.AccountNumber) == "" {
		return nil, errors.BadRequest("ACCOUNT_NUMBER_REQUIRED", "account_number is required")
	}

	account := mapper.NewAccountModel(customerID, req.AccountNumber, req.AccountType, req.Currency)

	if err := uc.repo.Create(ctx, account); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.Conflict("DUPLICATE_ACCOUNT_NUMBER", "account with this account_number already exists")
		}
		return nil, errors.InternalServer("CREATE_ACCOUNT_FAILED",err.Error())
	}
	return &pb.CreateAccountResponse{Account: mapper.MapAccountToProto(account)}, nil
}
