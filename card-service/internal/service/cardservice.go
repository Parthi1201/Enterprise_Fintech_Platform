package service

import (
	"context"


	pb "card-service/api/card"
	"card-service/internal/handler"

)

type CardServiceService struct {
    pb.UnimplementedCardServiceServer
    uc *handler.CardUsecase
}

func NewCardServiceService(uc *handler.CardUsecase) *CardServiceService {
    return &CardServiceService{uc: uc}
}

func (s *CardServiceService) CreateCard(ctx context.Context,req *pb.CreateCardRequest) (*pb.CreateCardResponse, error) {
	return s.uc.CreateCard(ctx, req)
}


func (s *CardServiceService) GetCard(ctx context.Context,req *pb.GetCardRequest) (*pb.Card, error) {
	return s.uc.GetCard(ctx, req)
}


func (s *CardServiceService) ListCardsByAccount(ctx context.Context,req *pb.ListCardsByAccountRequest) (*pb.ListCardsByAccountResponse, error) {
	return s.uc.ListCardsByAccount(ctx, req)
}


func (s *CardServiceService) UpdateCardStatus(ctx context.Context,req *pb.UpdateCardStatusRequest) (*pb.Card, error) {
	return s.uc.UpdateCardStatus(ctx, req)
}


func (s *CardServiceService) UpdateCardLimit(ctx context.Context,req *pb.UpdateCardLimitRequest) (*pb.Card, error) {
	return s.uc.UpdateCardLimit(ctx, req)
}
