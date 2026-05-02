package service

import (
	"context"

	pb "balance-service/api/balance"
	"balance-service/internal/handler"
	"balance-service/internal/mapper"

	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

type BalanceServiceService struct {
	pb.UnimplementedBalanceServiceServer
	uc *handler.BalanceUsecase
}

func NewBalanceServiceService(uc *handler.BalanceUsecase) *BalanceServiceService {
	return &BalanceServiceService{uc:uc}
}


func (s *BalanceServiceService) RecordBalanceChange(
	ctx context.Context,
	req *pb.RecordBalanceChangeRequest,
) (*pb.Balance, error) {

	balance, err := s.uc.RecordBalanceChange(
		ctx,
		req.AccountId,
		req.AvailableBalance,
		req.PendingBalance,
		req.Currency,
		req.ChangeAmount,
		req.ChangeReason,
	)
	if err != nil {
		return nil, errors.InternalServer("RECORD_BALANCE_FAILED", err.Error())
	}

	return mapper.MapBalanceToProto(balance), nil
}
func (s *BalanceServiceService) GetCurrentBalance(
	ctx context.Context,
	req *pb.GetCurrentBalanceRequest,
) (*pb.Balance, error) {

	balance, err := s.uc.GetCurrentBalance(ctx, req.AccountId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound("BALANCE_NOT_FOUND", "no balance found")
		}
		return nil, errors.InternalServer("GET_BALANCE_FAILED", err.Error())
	}

	return mapper.MapBalanceToProto(balance), nil
}

func (s *BalanceServiceService) GetBalanceHistory(
	ctx context.Context,
	req *pb.GetBalanceHistoryRequest,
) (*pb.GetBalanceHistoryResponse, error) {

	var fromDate *time.Time
	var toDate *time.Time

	if req.FromDate != "" {
		t, err := time.Parse("2006-01-02", req.FromDate)
		if err != nil {
			return nil, errors.BadRequest("INVALID_FROM_DATE", "use YYYY-MM-DD")
		}
		fromDate = &t
	}

	if req.ToDate != "" {
		t, err := time.Parse("2006-01-02", req.ToDate)
		if err != nil {
			return nil, errors.BadRequest("INVALID_TO_DATE", "use YYYY-MM-DD")
		}
		toDate = &t
	}

	balances, err := s.uc.GetBalanceHistory(
		ctx,
		req.AccountId,
		fromDate,
		toDate,
	)
	if err != nil {
		return nil, errors.InternalServer("GET_BALANCE_HISTORY_FAILED", err.Error())
	}

	resp := make([]*pb.Balance, 0, len(balances))
	for _, b := range balances {
		resp = append(resp, mapper.MapBalanceToProto(b))
	}

	return &pb.GetBalanceHistoryResponse{
		Balances: resp,
	}, nil
}
