package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/model"
)

type ProductQualityRepositoryMock struct {
	mock.Mock
}

func (mock *ProductQualityRepositoryMock) FindAll(ctx context.Context) ([]*model.ProductQuality, error) {
	args := mock.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*model.ProductQuality), args.Error(1)
}

func (mock *ProductQualityRepositoryMock) FindAllByProductCode(ctx context.Context, productCode string) ([]*model.ProductQuality, error) {
	args := mock.Called(ctx, productCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*model.ProductQuality), args.Error(1)
}

func (mock *ProductQualityRepositoryMock) FindByID(ctx context.Context, id int64) (*model.ProductQuality, error) {
	args := mock.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.ProductQuality), args.Error(1)
}

func (mock *ProductQualityRepositoryMock) Delete(ctx context.Context, id int64) error {
	args := mock.Called(ctx, id)
	return args.Error(0)
}
