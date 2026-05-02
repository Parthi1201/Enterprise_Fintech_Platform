package handler

import (
	"context"
	"strconv"

	pb "card-service/api/card"
	"card-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
)

func (uc *CardUsecase) ListCardsByAccount(ctx context.Context,req *pb.ListCardsByAccountRequest) (*pb.ListCardsByAccountResponse, error) {

	accountID, err := strconv.ParseInt(req.AccountId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest("INVALID_ACCOUNT_ID","account_id must be a valid numeric ID")
	}

	cards, err := uc.repo.ListByAccountID(ctx, accountID)
	if err != nil {
		return nil, errors.InternalServer("LIST_CARDS_FAILED",err.Error())
	}

	resp := make([]*pb.Card, 0, len(cards))
	for _, card := range cards {
		resp = append(resp, mapper.MapCardToProto(card))
	}

	return &pb.ListCardsByAccountResponse{Cards: resp}, nil
}
