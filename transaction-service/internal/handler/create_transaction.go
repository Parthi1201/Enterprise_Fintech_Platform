package handler

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"


	snowid "Enterprise_Fintech_Platform/common/id"
	pb "transaction-service/api/transaction"
	"transaction-service/internal/mapper"
)
func (uc *TransactionUsecase) CreateTransaction(
	ctx context.Context,
	req *pb.CreateTransactionRequest,
) (*pb.CreateTransactionResponse, error) {

	// account_id
	accountID, err := strconv.ParseInt(req.AccountId, 10, 64)
	if err != nil {
		return nil, errors.New("INVALID_ACCOUNT_ID")
	}

	// amount (float64 – because model forces it)
	amount, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil || amount <= 0 {
		return nil, errors.New("INVALID_AMOUNT")
	}

	// account existence
	if err := uc.accountClient.GetAccount(ctx, req.AccountId); err != nil {
		return nil, errors.New("ACCOUNT_NOT_FOUND")
	}

	// transaction type
	txType := strings.ToUpper(req.TransactionType)
	switch txType {
	case "DEBIT", "CREDIT", "REVERSAL":
	default:
		return nil, errors.New("INVALID_TRANSACTION_TYPE")
	}


	now := time.Now()

	tx := mapper.NewTransactionModel(
		int64(snowid.New()),
		accountID,
		txType,
		amount,
		req.Currency,
		req.Description,
		now,
	)

	if err := uc.repo.Create(ctx, tx); err != nil {
		return nil, err
	}

	return &pb.CreateTransactionResponse{
		Transaction: mapper.MapTransactionToProto(tx),
	}, nil
}
