package handler

import (
	"context"
	"strconv"

	pb "card-service/api/card"
	"card-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (uc *CardUsecase) GetCard(ctx context.Context,req *pb.GetCardRequest) (*pb.Card, error) {

	cardID, err := strconv.ParseInt(req.CardId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest("INVALID_CARD_ID","card_id must be a valid numeric ID")
	}

	card, err := uc.repo.GetByID(ctx, cardID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound("CARD_NOT_FOUND","card not found")
		}
		return nil, errors.InternalServer("GET_CARD_FAILED",err.Error())
	}
	return mapper.MapCardToProto(card), nil
}
