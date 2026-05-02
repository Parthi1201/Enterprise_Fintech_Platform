package service

import (
	"context"

	pb "account-service/api/account"
	"account-service/internal/handler"
)

type AccountServiceService struct {

	pb.UnimplementedAccountServiceServer
    uc *handler.AccountUsecase
}

func NewAccountServiceService(uc *handler.AccountUsecase) *AccountServiceService {
	return &AccountServiceService{uc:uc}
}

func (s *AccountServiceService) CreateAccount(ctx context.Context,req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	return s.uc.CreateAccount(ctx, req)
}

func (s *AccountServiceService) GetAccount(ctx context.Context,req *pb.GetAccountRequest) (*pb.Account, error){
	return s.uc.GetAccount(ctx, req)
}

func (s *AccountServiceService) UpdateAccountStatus(ctx context.Context,req *pb.UpdateAccountStatusRequest) (*pb.Account, error) {
	return s.uc.UpdateAccountStatus(ctx, req)
}

