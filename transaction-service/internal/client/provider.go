package client

import (
	pb "account-service/api/account"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/google/wire"
)

//wire for acc service
var ClientProviderSet = wire.NewSet(NewAccountServiceClient,NewAccountClient)


func NewAccountServiceClient()pb.AccountServiceClient {
		conn,err:=grpc.DialInsecure(context.Background(),grpc.WithEndpoint("127.0.0.1:9001")) // account grpc endpoint
	if err!=nil{
		panic(err)
	}

	return pb.NewAccountServiceClient(conn)
}

