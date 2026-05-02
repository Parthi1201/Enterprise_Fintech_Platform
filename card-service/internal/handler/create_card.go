package handler

import (
	"context"
	"strconv"
	"strings"

	pb "card-service/api/card"
	"card-service/internal/mapper"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (uc *CardUsecase) CreateCard(ctx context.Context,req *pb.CreateCardRequest) (*pb.CreateCardResponse, error) {

	accountID, err := strconv.ParseInt(req.AccountId, 10, 64)
	if err != nil {
		return nil, errors.BadRequest("INVALID_ACCOUNT_ID","account_id must be a valid numeric ID")
	}

	if err := uc.accountClient.GetAccount(
		ctx,
		strconv.FormatInt(accountID, 10),
	); err != nil {
		return nil, errors.NotFound("ACCOUNT_NOT_FOUND","account does not exist")
	}

	if strings.TrimSpace(req.CardNumber) == "" {
		return nil, errors.BadRequest("CARD_NUMBER_REQUIRED","card_number is required")
	}

	if strings.TrimSpace(req.CardType) == "" {
		return nil, errors.BadRequest("CARD_TYPE_REQUIRED","card_type is required")
	}

	card := mapper.NewCardModel(accountID,req.CardNumber,req.CardType)
	if err := uc.repo.Create(ctx, card); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.Conflict("DUPLICATE_CARD","card already exists")
		}
		return nil, errors.InternalServer("CREATE_CARD_FAILED",err.Error())
	}
	return &pb.CreateCardResponse{Card: mapper.MapCardToProto(card)}, nil
}
