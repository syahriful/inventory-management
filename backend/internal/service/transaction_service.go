package service

import (
	"context"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/repository"
)

type TransactionService struct {
	TransactionRepository    repository.TransactionRepositoryContract
	ProductQualityRepository repository.ProductQualityRepositoryContract
	TxTransactionRepository  repository.TxTransactionRepositoryContract
}

func NewTransactionService(transactionRepository repository.TransactionRepositoryContract, productQualityRepository repository.ProductQualityRepositoryContract, txTransactionRepository repository.TxTransactionRepositoryContract) TransactionServiceContract {
	return &TransactionService{
		TransactionRepository:    transactionRepository,
		ProductQualityRepository: productQualityRepository,
		TxTransactionRepository:  txTransactionRepository,
	}
}

func (service *TransactionService) FindAll(ctx context.Context) ([]*response.TransactionResponse, error) {
	transactions, err := service.TransactionRepository.FindAll(ctx, nil)
	if err != nil {
		return nil, err
	}

	var transactionResponses []*response.TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, transaction.ToResponse())
	}

	return transactionResponses, nil
}

func (service *TransactionService) FindAllBySupplierCode(ctx context.Context, supplierCode string) ([]*response.TransactionResponse, error) {
	transactions, err := service.TransactionRepository.FindAllBySupplierCode(ctx, supplierCode, nil)
	if err != nil {
		return nil, err
	}

	var transactionResponses []*response.TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, transaction.ToResponse())
	}

	return transactionResponses, nil
}

func (service *TransactionService) FindAllByCustomerCode(ctx context.Context, customerCode string) ([]*response.TransactionResponse, error) {
	transactions, err := service.TransactionRepository.FindAllByCustomerCode(ctx, customerCode, nil)
	if err != nil {
		return nil, err
	}

	var transactionResponses []*response.TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, transaction.ToResponse())
	}

	return transactionResponses, nil
}

func (service *TransactionService) FindByCode(ctx context.Context, code string) (*response.TransactionResponse, error) {
	transaction, err := service.TransactionRepository.FindByCodeWithAssociations(ctx, code, nil)
	if err != nil {
		return nil, err
	}

	return transaction.ToResponseWithAssociations(), nil
}

func (service *TransactionService) Create(ctx context.Context, request *request.CreateTransactionRequest) (*response.TransactionResponse, error) {
	transaction, err := service.TxTransactionRepository.Create(ctx, request)
	if err != nil {
		return nil, err
	}

	return transaction.ToResponse(), nil
}

func (service *TransactionService) Update(ctx context.Context, request *request.UpdateTransactionRequest) (*response.TransactionResponse, error) {
	transaction, err := service.TxTransactionRepository.Update(ctx, request)
	if err != nil {
		return nil, err
	}

	return transaction.ToResponse(), nil
}

func (service *TransactionService) TransferStock(ctx context.Context, request *request.TransferStockTransactionRequest) (*response.TransactionResponse, error) {
	transaction, err := service.TxTransactionRepository.TransferStock(ctx, request)
	if err != nil {
		return nil, err
	}

	return transaction.ToResponse(), nil
}

func (service *TransactionService) Delete(ctx context.Context, code string) error {
	err := service.TxTransactionRepository.Delete(ctx, code)
	if err != nil {
		return err
	}

	return nil
}
