package handler

import (
	"context"
	"strconv"
	"time"

	pb "payment-service/api/payment"
	"payment-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (uc *PaymentUsecase) CreatePayment(ctx context.Context,req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {

	fromAccountID, err := strconv.ParseInt(req.FromAccountId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest("INVALID_FROM_ACCOUNT_ID","from_account_id must be numeric")
	}

	if err := uc.accountClient.GetAccount(ctx,strconv.FormatInt(fromAccountID, 10)); err != nil {
		return nil, errors.NotFound("FROM_ACCOUNT_NOT_FOUND","from account does not exist")
	}

	var toAccountID int64
	if req.PaymentType == "transfer" {
		if req.ToAccountId == "" {
			return nil, errors.BadRequest("TO_ACCOUNT_REQUIRED","to_account_id is required for transfer")
		}
		toAccountID, err = strconv.ParseInt(req.ToAccountId, 10, 64)
		if err != nil {
			return nil, errors.BadRequest("INVALID_TO_ACCOUNT_ID","to_account_id must be numeric")
		}

		if err := uc.accountClient.GetAccount(ctx,strconv.FormatInt(toAccountID, 10)); err != nil {
			return nil, errors.NotFound("TO_ACCOUNT_NOT_FOUND","to account does not exist")
		}
	}

	amount, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil || amount <= 0 {
		return nil, errors.BadRequest("INVALID_AMOUNT","amount must be a positive number")
	}

	switch req.PaymentType {
	case "transfer", "payment", "topup":
	default:
		return nil, errors.BadRequest("INVALID_PAYMENT_TYPE","invalid payment_type")
	}
	switch req.PaymentMethod {
	case "card", "bank_transfer", "wallet":
	default:
		return nil, errors.BadRequest("INVALID_PAYMENT_METHOD","invalid payment_method")
	}

	var scheduledAt *time.Time
	if req.ScheduledAt != nil {
		t := req.ScheduledAt.AsTime()
		scheduledAt = &t
	}

	payment := mapper.NewPaymentModel(fromAccountID,toAccountID,req.PaymentType,req.PaymentMethod,amount,req.Currency,req.ReferenceNumber,req.ExternalReference,scheduledAt)

	if err := uc.repo.Create(ctx, payment); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.Conflict("DUPLICATE_REFERENCE","reference_number already exists")
		}
		return nil, errors.InternalServer("CREATE_PAYMENT_FAILED",err.Error())
	}

	return &pb.CreatePaymentResponse{Payment: mapper.MapPaymentToProto(payment)}, nil
}
