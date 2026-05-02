package client


import (
	"context"

	pb "account-service/api/account"
)

type AccountClient struct {
	client pb.AccountServiceClient
}

func NewAccountClient(client pb.AccountServiceClient) *AccountClient {
	return &AccountClient{client: client}
}

func (c *AccountClient) GetAccount(ctx context.Context,accountID string) error{
	_, err := c.client.GetAccount(ctx, &pb.GetAccountRequest{AccountId: accountID})
	return err
}
