package mapper

import (
	pb "customer-service/api/customer"
	"strconv"

	"customer-service/internal/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapCustomerToProto(customer *model.Customer) *pb.Customer {
	resp := &pb.Customer{
		CustomerId:     strconv.FormatInt(customer.CustomerID, 10),
		CustomerNumber: customer.CustomerNumber,
		FirstName:      customer.FirstName,
		LastName:       customer.LastName,
		Email:          customer.Email,
		Phone:          customer.Phone,
		Status:         customer.Status,
		KycStatus:      customer.KycStatus,
	}

	if customer.DateOfBirth != "" {
		resp.DateOfBirth = customer.DateOfBirth
	}
	if !customer.CreationTime.IsZero() {
		resp.CreationTime = timestamppb.New(customer.CreationTime)
	}
	if !customer.ModifiedTime.IsZero() {
		resp.ModifiedTime = timestamppb.New(customer.ModifiedTime)
	}

	return resp
}
