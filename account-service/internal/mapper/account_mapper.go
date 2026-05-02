package mapper

import (
	"strconv"

	pb "account-service/api/account"
	"account-service/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapAccountToProto(a *model.Account) *pb.Account {
	if a==nil {
		return nil
	}

	return &pb.Account{
		AccountId:strconv.FormatInt(a.AccountID,10),
		CustomerId:strconv.FormatInt(a.CustomerID,10),
		AccountNumber:a.AccountNumber,
		AccountType:a.AccountType,
		Currency:a.Currency,
		Status:a.Status,
		AvailableBalance:strconv.FormatFloat(a.AvailableBalance,'f',4,64),
		PendingBalance:strconv.FormatFloat(a.PendingBalance,'f',4,64),
		CreditLimit:strconv.FormatFloat(a.CreditLimit,'f',4,64),
		CreationTime:timestamppb.New(a.CreationTime),
		ModifiedTime:timestamppb.New(a.ModifiedTime),
	}
}
