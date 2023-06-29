package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
)

type TransactionServiceMock struct {
	mock.Mock
}

func (mock *TransactionServiceMock) FindAll(ctx context.Context) ([]*response.TransactionResponse, error) {
	args := mock.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*response.TransactionResponse), args.Error(1)
}

func (mock *TransactionServiceMock) FindAllBySupplierCode(ctx context.Context, supplierCode string) ([]*response.TransactionResponse, error) {
	args := mock.Called(ctx, supplierCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*response.TransactionResponse), args.Error(1)
}

func (mock *TransactionServiceMock) FindAllByCustomerCode(ctx context.Context, customerCode string) ([]*response.TransactionResponse, error) {
	args := mock.Called(ctx, customerCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*response.TransactionResponse), args.Error(1)
}

func (mock *TransactionServiceMock) FindByCode(ctx context.Context, code string) (*response.TransactionResponse, error) {
	args := mock.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.TransactionResponse), args.Error(1)
}

func (mock *TransactionServiceMock) Create(ctx context.Context, request *request.CreateTransactionRequest) (*response.TransactionResponse, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.TransactionResponse), args.Error(1)
}

func (mock *TransactionServiceMock) Update(ctx context.Context, request *request.UpdateTransactionRequest) (*response.TransactionResponse, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.TransactionResponse), args.Error(1)
}

func (mock *TransactionServiceMock) TransferStock(ctx context.Context, request *request.TransferStockTransactionRequest) (*response.TransactionResponse, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.TransactionResponse), args.Error(1)
}

func (mock *TransactionServiceMock) Delete(ctx context.Context, code string) error {
	args := mock.Called(ctx, code)
	return args.Error(0)
}
