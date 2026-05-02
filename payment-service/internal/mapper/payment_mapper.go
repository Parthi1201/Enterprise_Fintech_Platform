package mapper

import (
	"strconv"
	"time"

	snowid "Enterprise_Fintech_Platform/common/id"

	pb "payment-service/api/payment"
	"payment-service/internal/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewPaymentModel(fromAccountID int64,toAccountID int64,paymentType string,paymentMethod string,amount float64,currency string,referenceNumber string,externalReference string,scheduledAt *time.Time) *model.Payment {

	now := time.Now()
	id := snowid.New()

	p := &model.Payment{
		PaymentID:       int64(id),
		FromAccountID:   fromAccountID,
		PaymentType:     paymentType,
		PaymentMethod:   paymentMethod,
		Amount:          amount,
		Currency:        currency,
		Status:          "initiated",
		ReferenceNumber: referenceNumber,
		CreationTime:    now,
		ModifiedTime:    now,
	}

	if toAccountID != 0 {
		p.ToAccountID = toAccountID
	}

	if externalReference != "" {
		p.ExternalReference = externalReference
	}

	if scheduledAt != nil {
		p.ScheduledAt = *scheduledAt
	}

	return p
}

func MapPaymentToProto(m *model.Payment) *pb.Payment {
	if m == nil {
		return nil
	}

	p := &pb.Payment{
		PaymentId:       strconv.FormatInt(m.PaymentID, 10),
		FromAccountId:   strconv.FormatInt(m.FromAccountID, 10),
		PaymentType:     m.PaymentType,
		PaymentMethod:   m.PaymentMethod,
		Amount:          strconv.FormatFloat(m.Amount, 'f', -1, 64),
		Currency:        m.Currency,
		Status:          m.Status,
		ReferenceNumber: m.ReferenceNumber,
		ExternalReference: m.ExternalReference,
	}

	if m.ToAccountID != 0 {
		p.ToAccountId = strconv.FormatInt(m.ToAccountID, 10)
	}

	if !m.ScheduledAt.IsZero() {
		p.ScheduledAt = timestamppb.New(m.ScheduledAt)
	}

	if !m.ProcessedAt.IsZero() {
		p.ProcessedAt = timestamppb.New(m.ProcessedAt)
	}

	if !m.CreationTime.IsZero() {
		p.CreationTime = timestamppb.New(m.CreationTime)
	}

	if !m.ModifiedTime.IsZero() {
		p.ModifiedTime = timestamppb.New(m.ModifiedTime)
	}

	return p
}
