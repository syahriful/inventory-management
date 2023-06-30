package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/model"
)

type SupplierRepositoryMock struct {
	mock.Mock
}

func (mock *SupplierRepositoryMock) FindAll(ctx context.Context, offset int, limit int) ([]*model.Supplier, error) {
	args := mock.Called(ctx, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*model.Supplier), args.Error(1)
}

func (mock *SupplierRepositoryMock) CountAll(ctx context.Context) (int64, error) {
	args := mock.Called(ctx)
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}

	return args.Get(0).(int64), args.Error(1)
}

func (mock *SupplierRepositoryMock) FindByCodeWithAssociations(ctx context.Context, code string) (*model.Supplier, error) {
	args := mock.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Supplier), args.Error(1)
}

func (mock *SupplierRepositoryMock) FindByCode(ctx context.Context, code string) (*model.Supplier, error) {
	args := mock.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Supplier), args.Error(1)
}

func (mock *SupplierRepositoryMock) Create(ctx context.Context, supplier *model.Supplier) (*model.Supplier, error) {
	args := mock.Called(ctx, supplier)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Supplier), args.Error(1)
}

func (mock *SupplierRepositoryMock) Update(ctx context.Context, supplier *model.Supplier) (*model.Supplier, error) {
	args := mock.Called(ctx, supplier)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Supplier), args.Error(1)
}

func (mock *SupplierRepositoryMock) Delete(ctx context.Context, code string) error {
	args := mock.Called(ctx, code)
	return args.Error(0)
}
