package service

import (
	"context"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
	"inventory-management/backend/internal/model"
	"inventory-management/backend/internal/repository"
)

type TransactionService struct {
	TransactionRepository    repository.TransactionRepositoryContract
	ProductQualityRepository repository.ProductQualityRepositoryContract
	TxRepository             repository.TxTransactionRepositoryContract
}

func NewTransactionService(transactionRepository repository.TransactionRepositoryContract, productQualityRepository repository.ProductQualityRepositoryContract, txRepository repository.TxTransactionRepositoryContract) TransactionServiceContract {
	return &TransactionService{
		TransactionRepository:    transactionRepository,
		ProductQualityRepository: productQualityRepository,
		TxRepository:             txRepository,
	}
}

func (service *TransactionService) FindAll(ctx context.Context) ([]*response.TransactionResponse, error) {
	transactions, err := service.TransactionRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var transactionResponses []*response.TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, &response.TransactionResponse{
			ID:               transaction.ID,
			Code:             transaction.Code,
			ProductQualityID: transaction.ProductQualityID,
			SupplierCode:     transaction.SupplierCode,
			CustomerCode:     transaction.CustomerCode,
			Description:      transaction.Description,
			Quantity:         transaction.Quantity,
			Type:             transaction.Type,
			CreatedAt:        transaction.CreatedAt.String(),
		})
	}

	return transactionResponses, nil
}

func (service *TransactionService) FindAllBySupplierCode(ctx context.Context, supplierCode string) ([]*response.TransactionResponse, error) {
	transactions, err := service.TransactionRepository.FindAllBySupplierCode(ctx, supplierCode)
	if err != nil {
		return nil, err
	}

	var transactionResponses []*response.TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, &response.TransactionResponse{
			ID:               transaction.ID,
			Code:             transaction.Code,
			ProductQualityID: transaction.ProductQualityID,
			SupplierCode:     transaction.SupplierCode,
			CustomerCode:     transaction.CustomerCode,
			Description:      transaction.Description,
			Quantity:         transaction.Quantity,
			Type:             transaction.Type,
			CreatedAt:        transaction.CreatedAt.String(),
		})
	}

	return transactionResponses, nil
}

func (service *TransactionService) FindAllByCustomerCode(ctx context.Context, customerCode string) ([]*response.TransactionResponse, error) {
	transactions, err := service.TransactionRepository.FindAllByCustomerCode(ctx, customerCode)
	if err != nil {
		return nil, err
	}

	var transactionResponses []*response.TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, &response.TransactionResponse{
			ID:               transaction.ID,
			Code:             transaction.Code,
			ProductQualityID: transaction.ProductQualityID,
			SupplierCode:     transaction.SupplierCode,
			CustomerCode:     transaction.CustomerCode,
			Description:      transaction.Description,
			Quantity:         transaction.Quantity,
			Type:             transaction.Type,
			CreatedAt:        transaction.CreatedAt.String(),
		})
	}

	return transactionResponses, nil
}

func (service *TransactionService) FindByCode(ctx context.Context, code string) (*response.TransactionResponse, error) {
	transaction, err := service.TransactionRepository.FindByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	var supplierResponse *response.SupplierResponse
	if transaction.Supplier != nil {
		supplierResponse = &response.SupplierResponse{
			Code:    transaction.Supplier.Code,
			Name:    transaction.Supplier.Name,
			Address: transaction.Supplier.Address,
			Phone:   transaction.Supplier.Phone,
		}
	}

	var customerResponse *response.CustomerResponse
	if transaction.Customer != nil {
		customerResponse = &response.CustomerResponse{
			Code: transaction.Customer.Code,
			Name: transaction.Customer.Name,
		}
	}

	var productQualityResponse *response.ProductQualityResponse
	productQualityResponse = &response.ProductQualityResponse{
		ID:          transaction.ProductQuality.ID,
		ProductCode: transaction.ProductQuality.ProductCode,
		Quality:     transaction.ProductQuality.Quality,
		Price:       transaction.ProductQuality.Price,
		Quantity:    transaction.ProductQuality.Quantity,
		Type:        transaction.ProductQuality.Type,
		Product: &response.ProductResponse{
			Code:                transaction.ProductQuality.Product.Code,
			Name:                transaction.ProductQuality.Product.Name,
			UnitMassAcronym:     transaction.ProductQuality.Product.UnitMassAcronym,
			UnitMassDescription: transaction.ProductQuality.Product.UnitMassDescription,
		},
	}

	return &response.TransactionResponse{
		ID:               transaction.ID,
		Code:             transaction.Code,
		ProductQualityID: transaction.ProductQualityID,
		ProductQuality:   productQualityResponse,
		SupplierCode:     transaction.SupplierCode,
		Supplier:         supplierResponse,
		CustomerCode:     transaction.CustomerCode,
		Customer:         customerResponse,
		Description:      transaction.Description,
		Quantity:         transaction.Quantity,
		Type:             transaction.Type,
		CreatedAt:        transaction.CreatedAt.String(),
	}, nil
}

func (service *TransactionService) Create(ctx context.Context, request *request.CreateTransactionRequest) (*response.TransactionResponse, error) {
	transaction, err := service.TxRepository.Create(ctx, &model.Transaction{
		ProductQualityID: request.ProductQualityID,
		SupplierCode:     request.SupplierCode,
		CustomerCode:     request.CustomerCode,
		Description:      request.Description,
		Quantity:         request.Quantity,
		Type:             request.Type,
	})
	if err != nil {
		return nil, err
	}

	return &response.TransactionResponse{
		ID:               transaction.ID,
		Code:             transaction.Code,
		ProductQualityID: transaction.ProductQualityID,
		SupplierCode:     transaction.SupplierCode,
		CustomerCode:     transaction.CustomerCode,
		Description:      transaction.Description,
		Quantity:         transaction.Quantity,
		Type:             transaction.Type,
		CreatedAt:        transaction.CreatedAt.String(),
	}, nil
}

func (service *TransactionService) Update(ctx context.Context, request *request.UpdateTransactionRequest) (*response.TransactionResponse, error) {
	transaction, err := service.TransactionRepository.FindByCode(ctx, request.Code)
	if err != nil {
		return nil, err
	}

	transaction.SupplierCode = request.SupplierCode
	transaction.CustomerCode = request.CustomerCode
	transaction.Description = request.Description
	err = service.TxRepository.Update(ctx, request.Quantity, transaction)
	if err != nil {
		return nil, err
	}

	return &response.TransactionResponse{
		ID:               transaction.ID,
		Code:             transaction.Code,
		ProductQualityID: transaction.ProductQualityID,
		SupplierCode:     transaction.SupplierCode,
		CustomerCode:     transaction.CustomerCode,
		Description:      transaction.Description,
		Quantity:         transaction.Quantity,
		Type:             transaction.Type,
		CreatedAt:        transaction.CreatedAt.String(),
	}, nil
}

func (service *TransactionService) Delete(ctx context.Context, code string) error {
	transaction, err := service.TransactionRepository.FindByCode(ctx, code)
	if err != nil {
		return err
	}

	err = service.TxRepository.Delete(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}
