package handler

import (
	"context"
	"strconv"

	pb "payment-service/api/payment"
	"payment-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (uc *PaymentUsecase) GetPayment(ctx context.Context,req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error) {

	paymentID, err := strconv.ParseInt(req.PaymentId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest(
			"INVALID_PAYMENT_ID",
			"payment_id must be a valid numeric ID",
		)
	}

	payment, err := uc.repo.GetByID(ctx, paymentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound(
				"PAYMENT_NOT_FOUND",
				"payment not found",
			)
		}
		return nil, errors.InternalServer(
			"GET_PAYMENT_FAILED",
			err.Error(),
		)
	}

	return &pb.GetPaymentResponse{Payment: mapper.MapPaymentToProto(payment)}, nil
}
