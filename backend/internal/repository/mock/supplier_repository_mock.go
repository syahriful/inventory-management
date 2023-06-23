package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/model"
)

type SupplierRepositoryMock struct {
	mock.Mock
}

func (mock *SupplierRepositoryMock) FindAll(ctx context.Context) ([]*model.Supplier, error) {
	args := mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*model.Supplier), args.Error(1)
}

func (mock *SupplierRepositoryMock) FindByCodeWithAssociations(ctx context.Context, code string) (*model.Supplier, error) {
	args := mock.Called(code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Supplier), args.Error(1)
}

func (mock *SupplierRepositoryMock) FindByCode(ctx context.Context, code string) (*model.Supplier, error) {
	args := mock.Called(code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Supplier), args.Error(1)
}

func (mock *SupplierRepositoryMock) Create(ctx context.Context, supplier *model.Supplier) (*model.Supplier, error) {
	args := mock.Called(supplier)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Supplier), args.Error(1)
}

func (mock *SupplierRepositoryMock) Update(ctx context.Context, supplier *model.Supplier) (*model.Supplier, error) {
	args := mock.Called(supplier)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Supplier), args.Error(1)
}

func (mock *SupplierRepositoryMock) Delete(ctx context.Context, code string) error {
	args := mock.Called(code)
	return args.Error(0)
}
