package service

import (
	"context"

	pb "transaction-service/api/transaction"
	"transaction-service/internal/handler"
)

type TransactionServiceService struct {
	pb.UnimplementedTransactionServiceServer
	
	uc *handler.TransactionUsecase
}

func NewTransactionServiceService(uc *handler.TransactionUsecase) *TransactionServiceService {
	return &TransactionServiceService{uc: uc}
}

func (s *TransactionServiceService) CreateTransaction(ctx context.Context,req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	return s.uc.CreateTransaction(ctx, req)
}

func (s *TransactionServiceService) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.Transaction, error) {
    return &pb.Transaction{}, nil
}
func (s *TransactionServiceService) ListTransactionsByAccount(ctx context.Context, req *pb.ListTransactionsByAccountRequest) (*pb.ListTransactionsByAccountResponse, error) {
    return &pb.ListTransactionsByAccountResponse{}, nil
}
