package service

import (
	"context"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/model"
	"inventory-management/backend/internal/repository"
	"inventory-management/backend/util"
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
		transactionResponses = append(transactionResponses, transaction.ToResponse())
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
		transactionResponses = append(transactionResponses, transaction.ToResponse())
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
		transactionResponses = append(transactionResponses, transaction.ToResponse())
	}

	return transactionResponses, nil
}

func (service *TransactionService) FindByCode(ctx context.Context, code string) (*response.TransactionResponse, error) {
	transaction, err := service.TransactionRepository.FindByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	return transaction.ToResponseWithAssociations(), nil
}

func (service *TransactionService) Create(ctx context.Context, request *request.CreateTransactionRequest) (*response.TransactionResponse, error) {
	productQuality, err := service.ProductQualityRepository.FindByIDWithAssociations(ctx, request.ProductQualityID)
	if err != nil {
		return nil, err
	}

	quantity, err := util.CalculateUnitOfMass(productQuality.Product.UnitMassAcronym, request.UnitMassAcronym, request.Quantity)
	if err != nil {
		return nil, err
	}

	var transactionRequest model.Transaction
	transactionRequest.ProductQualityID = request.ProductQualityID
	transactionRequest.ProductQualityIDTransferred = request.ProductQualityIDTransferred
	transactionRequest.SupplierCode = request.SupplierCode
	transactionRequest.CustomerCode = request.CustomerCode
	transactionRequest.Description = request.Description
	transactionRequest.Quantity = request.Quantity
	transactionRequest.Type = request.Type
	transactionRequest.UnitMassAcronym = request.UnitMassAcronym

	transaction, err := service.TxRepository.Create(ctx, &transactionRequest, quantity)
	if err != nil {
		return nil, err
	}

	return transaction.ToResponse(), nil
}

func (service *TransactionService) Update(ctx context.Context, request *request.UpdateTransactionRequest) (*response.TransactionResponse, error) {
	transaction, err := service.TransactionRepository.FindByCode(ctx, request.Code)
	if err != nil {
		return nil, err
	}

	transactionQuantity, err := util.CalculateUnitOfMass(transaction.ProductQuality.Product.UnitMassAcronym, transaction.UnitMassAcronym, transaction.Quantity)
	if err != nil {
		return nil, err
	}

	if request.CustomerCode != nil {
		transaction.CustomerCode = request.CustomerCode
	}

	if request.SupplierCode != nil {
		transaction.SupplierCode = request.SupplierCode
	}

	transaction.Description = request.Description
	transaction.UnitMassAcronym = request.UnitMassAcronym
	transaction.Quantity = transactionQuantity

	increaseStock, err := util.CalculateUnitOfMass(transaction.ProductQuality.Product.UnitMassAcronym, request.UnitMassAcronym, request.Quantity)
	if err != nil {
		return nil, err
	}

	err = service.TxRepository.Update(ctx, increaseStock, request.Quantity, transaction)
	if err != nil {
		return nil, err
	}

	return transaction.ToResponse(), nil
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
