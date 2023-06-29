package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/model"
)

type TxTransactionRepositoryMock struct {
	mock.Mock
}

func (mock *TxTransactionRepositoryMock) Create(ctx context.Context, request *request.CreateTransactionRequest) (*model.Transaction, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (mock *TxTransactionRepositoryMock) TransferStock(ctx context.Context, request *request.TransferStockTransactionRequest) (*model.Transaction, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (mock *TxTransactionRepositoryMock) Update(ctx context.Context, request *request.UpdateTransactionRequest) (*model.Transaction, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (mock *TxTransactionRepositoryMock) Delete(ctx context.Context, code string) error {
	args := mock.Called(ctx, code)
	return args.Error(0)
}
