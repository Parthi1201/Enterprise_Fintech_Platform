package client

import (
	"context"

	pb "customer-service/api/customer"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewCustomerServiceClient()pb.CustomerServiceClient {
	conn,err:=grpc.DialInsecure(context.Background(),grpc.WithEndpoint("127.0.0.1:9000")) // customer grpc endpoint
	if err!=nil{
		panic(err)
	}

	return pb.NewCustomerServiceClient(conn)
}


