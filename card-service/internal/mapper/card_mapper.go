package mapper

import (
	"strconv"

	pb "card-service/api/card"
	"card-service/internal/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapCardToProto(c *model.Card) *pb.Card {
	if c==nil{
		return nil
	}

	return &pb.Card{
		CardId:strconv.FormatInt(c.CardID, 10),
		AccountId:strconv.FormatInt(c.AccountID, 10),
		CardType:c.CardType,
		CardStatus:c.CardStatus,
		ExpiryDate:timestamppb.New(c.ExpiryDate),
		DailyLimit:"0",
		MonthlyLimit:"0",
		CreationTime:timestamppb.New(c.CreationTime),
		ModifiedTime:timestamppb.New(c.ModifiedTime),
	}
}
