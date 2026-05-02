package service

import (
	"context"
	pb "payment-service/api/payment"
	"payment-service/internal/handler"
)

type PaymentServiceService struct {
	pb.UnimplementedPaymentServiceServer

	uc *handler.PaymentUsecase
}

func NewPaymentServiceService(uc *handler.PaymentUsecase) *PaymentServiceService {
	return &PaymentServiceService{uc: uc}
}

func (s *PaymentServiceService) CreatePayment(ctx context.Context,req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	return s.uc.CreatePayment(ctx, req)
}

func (s *PaymentServiceService) GetPayment(ctx context.Context,req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error) {
	return s.uc.GetPayment(ctx, req)
}

func (s *PaymentServiceService) ListPaymentsByAccount(ctx context.Context,req *pb.ListPaymentsByAccountRequest) (*pb.ListPaymentsByAccountResponse, error) {
	return s.uc.ListPaymentsByAccount(ctx, req)
}

func (s *PaymentServiceService) UpdatePaymentStatus(ctx context.Context,req *pb.UpdatePaymentStatusRequest)(*pb.Payment, error) {
	return s.uc.UpdatePaymentStatus(ctx, req)
}

