package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"inventory-management/backend/internal/model"
)

type ProductQualityRepositoryMock struct {
	mock.Mock
}

func (mock *ProductQualityRepositoryMock) IncreaseStock(ctx context.Context, id int64, quantity float64, tx *gorm.DB) error {
	args := mock.Called(ctx, id, quantity)
	return args.Error(1)
}

func (mock *ProductQualityRepositoryMock) DecreaseStock(ctx context.Context, id int64, quantity float64, tx *gorm.DB) error {
	args := mock.Called(ctx, id, quantity)
	return args.Error(1)
}

func (mock *ProductQualityRepositoryMock) FindAll(ctx context.Context, tx *gorm.DB) ([]*model.ProductQuality, error) {
	args := mock.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*model.ProductQuality), args.Error(1)
}

func (mock *ProductQualityRepositoryMock) FindAllByProductCode(ctx context.Context, productCode string, tx *gorm.DB) ([]*model.ProductQuality, error) {
	args := mock.Called(ctx, productCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*model.ProductQuality), args.Error(1)
}

func (mock *ProductQualityRepositoryMock) FindByID(ctx context.Context, id int64, tx *gorm.DB) (*model.ProductQuality, error) {
	args := mock.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.ProductQuality), args.Error(1)
}

func (mock *ProductQualityRepositoryMock) FindByIDWithAssociations(ctx context.Context, id int64, tx *gorm.DB) (*model.ProductQuality, error) {
	args := mock.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.ProductQuality), args.Error(1)
}

func (mock *ProductQualityRepositoryMock) Delete(ctx context.Context, id int64, tx *gorm.DB) error {
	args := mock.Called(ctx, id)
	return args.Error(0)
}
