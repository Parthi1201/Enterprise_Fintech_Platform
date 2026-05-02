package service

import (
	"context"
	// "strconv"

	pb "customer-service/api/customer"

	"customer-service/internal/handler"
	// "customer-service/internal/mapper"
	// "github.com/go-kratos/kratos/v2/errors"
	// "google.golang.org/protobuf/types/known/timestamppb"
)

type CustomerServiceService struct {
	pb.UnimplementedCustomerServiceServer
	uc *handler.CustomerUseCase
}

func NewCustomerServiceService(uc *handler.CustomerUseCase) *CustomerServiceService {
	return &CustomerServiceService{uc: uc}
}

// create
func (s *CustomerServiceService) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	return s.uc.CreateCustomer(ctx, req)
}

// get
func (s *CustomerServiceService) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	return s.uc.GetCustomer(ctx, req)
}

// update(patch)
func (s *CustomerServiceService) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerRequest) (*pb.UpdateCustomerResponse, error) {
	return s.uc.UpdateCustomer(ctx, req)
}

func (s *CustomerServiceService) UpdateKycStatus(ctx context.Context, req *pb.UpdateKycStatusRequest) (*pb.Customer, error) {
	return s.uc.UpdateKycStatus(ctx, req)
}

// update status
func (s *CustomerServiceService) UpdateCustomerStatus(ctx context.Context, req *pb.UpdateCustomerStatusRequest) (*pb.Customer, error) {
	return s.uc.UpdateCustomerStatus(ctx, req)
}

// discard

func (s *CustomerServiceService) DiscardCustomer(ctx context.Context, req *pb.DiscardCustomerRequest) (*pb.Customer, error) {
	return s.uc.DiscardCustomer(ctx, req)
}
