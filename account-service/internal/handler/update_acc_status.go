package handler

import (
	"context"
	"strconv"

	pb "account-service/api/account"
	"account-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (uc *AccountUsecase) UpdateAccountStatus(ctx context.Context,req *pb.UpdateAccountStatusRequest)(*pb.Account, error) {

	accountID, err := strconv.ParseInt(req.AccountId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest("INVALID_ACCOUNT_ID","account_id must be a valid numeric ID")
	}

	switch req.Status {
	case "active", "frozen", "closed":
	default:
		return nil, errors.BadRequest("INVALID_ACCOUNT_STATUS","status must be active, frozen, or closed")
	}

	if err := uc.repo.UpdateStatus(ctx, accountID, req.Status); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound("ACCOUNT_NOT_FOUND","account not found")
		}
		return nil, errors.InternalServer("UPDATE_STATUS_FAILED",err.Error())
	}

	account, err := uc.repo.GetByID(ctx, accountID)
	if err != nil {
		return nil, errors.InternalServer("FETCH_ACCOUNT_FAILED",err.Error())
	}
	return mapper.MapAccountToProto(account), nil
}
