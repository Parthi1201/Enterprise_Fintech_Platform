package handler

import (
	"context"
	"strconv"

	pb "card-service/api/card"
	"card-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (uc *CardUsecase) UpdateCardLimit(ctx context.Context,req *pb.UpdateCardLimitRequest) (*pb.Card, error) {

	cardID, err := strconv.ParseInt(req.CardId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest(
			"INVALID_CARD_ID",
			"card_id must be a valid numeric ID",
		)
	}

	dailyLimit, err := strconv.ParseFloat(req.DailyLimit, 64)
	if err != nil {
		return nil, errors.BadRequest(
			"INVALID_DAILY_LIMIT",
			"daily_limit must be numeric",
		)
	}

	monthlyLimit, err := strconv.ParseFloat(req.MonthlyLimit, 64)
	if err != nil {
		return nil, errors.BadRequest(
			"INVALID_MONTHLY_LIMIT",
			"monthly_limit must be numeric",
		)
	}

	if err := uc.repo.UpdateLimit(ctx, cardID, dailyLimit, monthlyLimit); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound(
				"CARD_NOT_FOUND",
				"card not found",
			)
		}
		return nil, errors.InternalServer(
			"UPDATE_CARD_LIMIT_FAILED",
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
