package service

import (
	"context"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/model"
	"inventory-management/backend/internal/repository"
)

type CustomerService struct {
	CustomerRepository repository.CustomerRepositoryContract
}

func NewCustomerService(customerRepository repository.CustomerRepositoryContract) CustomerServiceContract {
	return &CustomerService{
		CustomerRepository: customerRepository,
	}
}

func (service *CustomerService) FindAll(ctx context.Context) ([]*response.CustomerResponse, error) {
	customers, err := service.CustomerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var customerResponses []*response.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, customer.ToResponse())
	}

	return customerResponses, nil
}

func (service *CustomerService) FindByCode(ctx context.Context, code string) (*response.CustomerResponse, error) {
	customer, err := service.CustomerRepository.FindByCodeWithAssociations(ctx, code)
	if err != nil {
		return nil, err
	}

	return customer.ToResponseWithAssociations(), nil
}

func (service *CustomerService) Create(ctx context.Context, request *request.CreateCustomerRequest) (*response.CustomerResponse, error) {
	var customerRequest model.Customer
	customerRequest.Name = request.Name

	customer, err := service.CustomerRepository.Create(ctx, &customerRequest)
	if err != nil {
		return nil, err
	}

	return customer.ToResponse(), nil
}

func (service *CustomerService) Update(ctx context.Context, request *request.UpdateCustomerRequest) (*response.CustomerResponse, error) {
	checkCustomer, err := service.CustomerRepository.FindByCode(ctx, request.Code)
	if err != nil {
		return nil, err
	}

	checkCustomer.Name = request.Name
	customer, err := service.CustomerRepository.Update(ctx, checkCustomer)
	if err != nil {
		return nil, err
	}

	return customer.ToResponse(), nil
}

func (service *CustomerService) Delete(ctx context.Context, code string) error {
	checkCustomer, err := service.CustomerRepository.FindByCode(ctx, code)
	if err != nil {
		return err
	}

	err = service.CustomerRepository.Delete(ctx, checkCustomer.Code)
	if err != nil {
		return err
	}

	return nil
}
