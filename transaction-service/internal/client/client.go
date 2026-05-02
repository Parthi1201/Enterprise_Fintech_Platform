package client

import (
	"context"

	pb "account-service/api/account"
)

type AccountClient struct {
	client pb.AccountServiceClient
}

func NewAccountClient(grpcClient pb.AccountServiceClient)*AccountClient {
	return &AccountClient{client: grpcClient}
}

func(c *AccountClient) GetAccount(ctx context.Context,AccountID string) error {
	_,err:=c.client.GetAccount(ctx,&pb.GetAccountRequest{AccountId: AccountID})
	return err
}
