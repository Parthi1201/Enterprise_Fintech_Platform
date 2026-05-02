package handler

import (
	"context"
	"strconv"

	pb "payment-service/api/payment"
	"payment-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
)

func (uc *PaymentUsecase) ListPaymentsByAccount(ctx context.Context,req *pb.ListPaymentsByAccountRequest) (*pb.ListPaymentsByAccountResponse, error) {

	accountID, err := strconv.ParseInt(req.AccountId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest(
			"INVALID_ACCOUNT_ID",
			"account_id must be a valid numeric ID",
		)
	}

	if err := uc.accountClient.GetAccount(ctx, req.AccountId); err != nil {
		return nil, errors.NotFound(
			"ACCOUNT_NOT_FOUND",
			"account does not exist",
		)
	}

	payments, err := uc.repo.ListByAccountID(ctx, accountID)
	if err != nil {
		return nil, errors.InternalServer(
			"LIST_PAYMENTS_FAILED",
			err.Error(),
		)
	}

	resp := make([]*pb.Payment, 0, len(payments))
	for _, p := range payments {
		resp = append(resp, mapper.MapPaymentToProto(p))
	}

	return &pb.ListPaymentsByAccountResponse{Payments: resp}, nil
}
