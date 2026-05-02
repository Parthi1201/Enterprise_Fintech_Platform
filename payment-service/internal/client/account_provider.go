package account

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ProviderSet = wire.NewSet(
	NewClient,
	NewGRPCConn,
)

// grpc connection to account-service
func NewGRPCConn() (*grpc.ClientConn, error) {
	return grpc.Dial(
		"127.0.0.1:9001", // account-service grpc port
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}
