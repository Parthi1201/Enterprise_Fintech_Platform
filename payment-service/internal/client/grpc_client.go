package account

import (
	"context"

	pb "account-service/api/account"
	"google.golang.org/grpc"
)

type Client struct {
	client pb.AccountServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		client: pb.NewAccountServiceClient(conn),
	}
}

func (c *Client) GetAccount(ctx context.Context, accountID string) error {
	_, err := c.client.GetAccount(ctx, &pb.GetAccountRequest{
		AccountId: accountID,
	})
	return err
}
