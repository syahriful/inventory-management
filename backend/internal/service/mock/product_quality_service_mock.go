package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/http/response"
)

type ProductQualityServiceMock struct {
	mock.Mock
}

func (mock *ProductQualityServiceMock) FindAll(ctx context.Context) ([]*response.ProductQualityResponse, error) {
	args := mock.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*response.ProductQualityResponse), args.Error(1)
}

func (mock *ProductQualityServiceMock) FindAllByProductCode(ctx context.Context, productCode string) (*response.ProductQualityWithOwnProductResponse, error) {
	args := mock.Called(ctx, productCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.ProductQualityWithOwnProductResponse), args.Error(1)
}

func (mock *ProductQualityServiceMock) FindByID(ctx context.Context, id int64) (*response.ProductQualityResponse, error) {
	args := mock.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.ProductQualityResponse), args.Error(1)
}

func (mock *ProductQualityServiceMock) Delete(ctx context.Context, id int64) error {
	args := mock.Called(ctx, id)
	return args.Error(0)
}
