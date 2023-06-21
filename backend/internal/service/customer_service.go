package service

import (
	"context"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
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
		customerResponses = append(customerResponses, &response.CustomerResponse{
			ID:        customer.ID,
			Code:      customer.Code,
			Name:      customer.Name,
			CreatedAt: customer.CreatedAt.String(),
			UpdatedAt: customer.UpdatedAt.String(),
		})
	}

	return customerResponses, nil
}

func (service *CustomerService) FindByCode(ctx context.Context, code string) (*response.CustomerResponse, error) {
	customer, err := service.CustomerRepository.FindByCodeWithAssociations(ctx, code)
	if err != nil {
		return nil, err
	}

	var transactionResponses []*response.TransactionResponse
	for _, transaction := range customer.Transactions {
		var productQualityResponse response.ProductQualityResponse
		productQualityResponse.ID = transaction.ProductQuality.ID
		productQualityResponse.Quality = transaction.ProductQuality.Quality
		productQualityResponse.Price = transaction.ProductQuality.Price
		productQualityResponse.Product = &response.ProductResponse{
			Code: transaction.ProductQuality.Product.Code,
			Name: transaction.ProductQuality.Product.Name,
		}

		transactionResponses = append(transactionResponses, &response.TransactionResponse{
			ID:               transaction.ID,
			Code:             transaction.Code,
			ProductQuality:   &productQualityResponse,
			ProductQualityID: transaction.ProductQualityID,
			CustomerCode:     transaction.CustomerCode,
			Description:      transaction.Description,
			Quantity:         transaction.Quantity,
			Type:             transaction.Type,
			CreatedAt:        transaction.CreatedAt.String(),
		})
	}

	return &response.CustomerResponse{
		ID:           customer.ID,
		Code:         customer.Code,
		Name:         customer.Name,
		CreatedAt:    customer.CreatedAt.String(),
		UpdatedAt:    customer.UpdatedAt.String(),
		Transactions: transactionResponses,
	}, nil
}

func (service *CustomerService) Create(ctx context.Context, request *request.CreateCustomerRequest) (*response.CustomerResponse, error) {
	customer, err := service.CustomerRepository.Create(ctx, &model.Customer{
		Name: request.Name,
	})
	if err != nil {
		return nil, err
	}

	return &response.CustomerResponse{
		ID:        customer.ID,
		Code:      customer.Code,
		Name:      customer.Name,
		CreatedAt: customer.CreatedAt.String(),
		UpdatedAt: customer.UpdatedAt.String(),
	}, nil
}

func (service *CustomerService) Update(ctx context.Context, request *request.UpdateCustomerRequest) (*response.CustomerResponse, error) {
	checkCustomer, err := service.CustomerRepository.FindByCode(ctx, request.Code)
	if err != nil {
		return nil, err
	}

	customer, err := service.CustomerRepository.Update(ctx, &model.Customer{
		Code: checkCustomer.Code,
		Name: request.Name,
	})
	if err != nil {
		return nil, err
	}

	return &response.CustomerResponse{
		ID:        customer.ID,
		Code:      customer.Code,
		Name:      customer.Name,
		CreatedAt: customer.CreatedAt.String(),
		UpdatedAt: customer.UpdatedAt.String(),
	}, nil
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
