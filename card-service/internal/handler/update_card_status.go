package handler

import (
	"context"
	"strconv"

	pb "card-service/api/card"
	"card-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

var validCardStatus = map[string]bool{
	"active":  true,
	"blocked": true,
	"expired": true,
}

func (uc *CardUsecase) UpdateCardStatus(ctx context.Context,req *pb.UpdateCardStatusRequest) (*pb.Card, error) {

	cardID, err := strconv.ParseInt(req.CardId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest(
			"INVALID_CARD_ID",
			"card_id must be a valid numeric ID",
		)
	}

	if !validCardStatus[req.CardStatus] {
		return nil, errors.BadRequest(
			"INVALID_CARD_STATUS",
			"status must be active, blocked, or expired",
		)
	}

	if err := uc.repo.UpdateStatus(ctx, cardID, req.CardStatus); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound(
				"CARD_NOT_FOUND",
				"card not found",
			)
		}
		return nil, errors.InternalServer(
			"UPDATE_CARD_STATUS_FAILED",
			err.Error(),
		)
	}

	card, err := uc.repo.GetByID(ctx, cardID)
	if err != nil {
		return nil, errors.InternalServer(
			"FETCH_CARD_FAILED",
			err.Error(),
		)
	}
	return mapper.MapCardToProto(card), nil
}
