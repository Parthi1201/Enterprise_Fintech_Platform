package mapper

import (
	snowid "Enterprise_Fintech_Platform/common/id"

	pb "customer-service/api/customer"

	"customer-service/internal/model"
)

func HandlerMapper(req *pb.CreateCustomerRequest)*model.Customer{
	id:=snowid.New()

	customer := &model.Customer{
		CustomerID:     int64(id),
		CustomerNumber: req.CustomerNumber,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		DateOfBirth:    req.DateOfBirth,
		Status:         "active",
		KycStatus:      "pending",
	}
	return customer
}
