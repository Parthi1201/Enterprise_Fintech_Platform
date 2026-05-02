package client

import (
	"context"

	pb "account-service/api/account"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewAccountServiceClient() pb.AccountServiceClient {
	conn, err := grpc.DialInsecure(context.Background(),grpc.WithEndpoint("127.0.0.1:9001")) // account service grpc port)
	if err != nil {
		panic(err)
	}

	return pb.NewAccountServiceClient(conn)
}
