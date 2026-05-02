package mapper

import (
	"strconv"
	"strings"
	"time"

	pb "transaction-service/api/transaction"
	"transaction-service/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapTransactionToProto(tx *model.Transaction) *pb.Transaction {
	if tx == nil {
		return nil
	}

	res := &pb.Transaction{
		TransactionId:   strconv.FormatInt(tx.TransactionID, 10),
		AccountId:       strconv.FormatInt(tx.AccountID, 10),
		TransactionType: strings.ToUpper(tx.TransactionType),
		Amount:          strconv.FormatFloat(tx.Amount, 'f', -1, 64),
		Currency:        tx.Currency,
		Status:          tx.Status,
		ReferenceNumber: tx.ReferenceNumber,
		MerchantInfo:    tx.MerchantInfo,
		Description:     tx.Description,
		TransactionTime: timestamppb.New(tx.TransactionTime),
	}

	// posted_at only if set
	if !tx.PostedAt.Equal(time.Time{}) {
		res.PostedAt = timestamppb.New(tx.PostedAt)
	}

	return res
}
