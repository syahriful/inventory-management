package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/model"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (mock *ProductRepositoryMock) FindAll(ctx context.Context, offset int, limit int) ([]*model.Product, error) {
	args := mock.Called(ctx, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*model.Product), args.Error(1)
}

func (mock *ProductRepositoryMock) CountAll(ctx context.Context) (int64, error) {
	args := mock.Called(ctx)
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}

	return args.Get(0).(int64), args.Error(1)
}

func (mock *ProductRepositoryMock) FindByCodeWithAssociations(ctx context.Context, code string) (*model.Product, error) {
	args := mock.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Product), args.Error(1)
}

func (mock *ProductRepositoryMock) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	args := mock.Called(ctx, product)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Product), args.Error(1)
}

func (mock *ProductRepositoryMock) Update(ctx context.Context, product *model.Product) (*model.Product, error) {
	args := mock.Called(ctx, product)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Product), args.Error(1)
}

func (mock *ProductRepositoryMock) Delete(ctx context.Context, code string) error {
	args := mock.Called(ctx, code)
	return args.Error(0)
}
