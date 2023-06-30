package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"inventory-management/backend/internal/model"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (mock *TransactionRepositoryMock) FindAll(ctx context.Context, offset int, limit int, tx *gorm.DB) ([]*model.Transaction, error) {
	args := mock.Called(ctx, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*model.Transaction), args.Error(1)
}

func (mock *TransactionRepositoryMock) CountAll(ctx context.Context, tx *gorm.DB) (int64, error) {
	args := mock.Called(ctx)
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}

	return args.Get(0).(int64), args.Error(1)
}

func (mock *TransactionRepositoryMock) FindAllBySupplierCode(ctx context.Context, supplierCode string, tx *gorm.DB) ([]*model.Transaction, error) {
	args := mock.Called(ctx, supplierCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*model.Transaction), args.Error(1)
}

func (mock *TransactionRepositoryMock) FindAllByCustomerCode(ctx context.Context, customerCode string, tx *gorm.DB) ([]*model.Transaction, error) {
	args := mock.Called(ctx, customerCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*model.Transaction), args.Error(1)
}

func (mock *TransactionRepositoryMock) FindByCodeWithAssociations(ctx context.Context, code string, tx *gorm.DB) (*model.Transaction, error) {
	args := mock.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (mock *TransactionRepositoryMock) FindByCode(ctx context.Context, code string, tx *gorm.DB) (*model.Transaction, error) {
	args := mock.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (mock *TransactionRepositoryMock) Create(ctx context.Context, transaction *model.Transaction, tx *gorm.DB) (*model.Transaction, error) {
	args := mock.Called(ctx, transaction)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (mock *TransactionRepositoryMock) Update(ctx context.Context, transaction *model.Transaction, tx *gorm.DB) (*model.Transaction, error) {
	args := mock.Called(ctx, transaction)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (mock *TransactionRepositoryMock) Delete(ctx context.Context, code string, tx *gorm.DB) error {
	args := mock.Called(ctx, code)
	return args.Error(1)
}
