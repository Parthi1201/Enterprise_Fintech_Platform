package mapper

import (
	pb "balance-service/api/balance"
	"balance-service/internal/data"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapBalanceToProto(b *data.Balance) *pb.Balance {
	return &pb.Balance{
		AccountId:b.AccountID,
		AvailableBalance:b.AvailableBalance,
		PendingBalance:b.PendingBalance,
		Currency:b.Currency,
		ChangeAmount:b.ChangeAmount,
		ChangeReason:b.ChangeReason,
		BalanceTimestamp:timestamppb.New(b.BalanceTimestamp),
		BalanceDate:timestamppb.New(b.BalanceDate),
	}
}
