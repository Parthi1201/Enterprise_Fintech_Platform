package client

import (
	"context"

	pb "customer-service/api/customer"
)

type CustomerClient struct{
	client pb.CustomerServiceClient
}

func NewCustomerClient(grpcClient pb.CustomerServiceClient)*CustomerClient {
	return &CustomerClient{client: grpcClient}
}


func(c *CustomerClient) GetCustomer(ctx context.Context,customerID string) error {
	_,err:=c.client.GetCustomer(ctx,&pb.GetCustomerRequest{CustomerId: customerID})
	return err
}
