package handler

import "customer-service/internal/data"

type CustomerUseCase struct{
	repo *data.CustomerRepo
}

func NewCustomerUseCase(repo *data.CustomerRepo)*CustomerUseCase{
	return &CustomerUseCase{repo:repo}
}

