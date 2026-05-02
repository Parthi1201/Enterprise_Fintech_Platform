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

var validTransitions = map[string]map[string]bool{
	"initiated": {
		"processing": true,
		"failed":     true,
	},
	"processing": {
		"completed": true,
		"failed":    true,
	},
}

func (uc *PaymentUsecase) UpdatePaymentStatus(ctx context.Context,req *pb.UpdatePaymentStatusRequest) (*pb.Payment, error) {

	paymentID, err := strconv.ParseInt(req.PaymentId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest("INVALID_PAYMENT_ID","payment_id must be a valid numeric ID")
	}

	payment, err := uc.repo.GetByID(ctx, paymentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound("PAYMENT_NOT_FOUND","payment not found")
		}
		return nil, errors.InternalServer("GET_PAYMENT_FAILED",err.Error())
	}

	if !validTransitions[payment.Status][req.Status] {
		return nil, errors.BadRequest("INVALID_PAYMENT_STATUS_TRANSITION","invalid payment status transition")
	}

	payment.Status = req.Status
	if req.Status == "completed" || req.Status == "failed" {
		payment.ProcessedAt = time.Now()
	}

	if err := uc.repo.Update(ctx, payment); err != nil {
		return nil, errors.InternalServer(
			"UPDATE_PAYMENT_STATUS_FAILED",
			err.Error(),
		)
	}

	return mapper.MapPaymentToProto(payment), nil
}
