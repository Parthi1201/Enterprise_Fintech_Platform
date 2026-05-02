package handler

import (
	"context"
	"strconv"

	pb "account-service/api/account"
	"account-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (uc *AccountUsecase) GetAccount(ctx context.Context,req *pb.GetAccountRequest) (*pb.Account, error) {

	accountID, err := strconv.ParseInt(req.AccountId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest("INVALID_ACCOUNT_ID","account_id must be a valid numeric ID")
	}
	account, err := uc.repo.GetByID(ctx, accountID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound("ACCOUNT_NOT_FOUND","account not found")
		}
		return nil, errors.InternalServer("GET_ACCOUNT_FAILED",err.Error())
	}
	return mapper.MapAccountToProto(account), nil
}
