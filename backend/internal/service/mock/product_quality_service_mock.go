package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/http/presenter/response"
)

type ProductQualityMock struct {
	mock.Mock
}

func (m *ProductQualityMock) FindAll(ctx context.Context) ([]*response.ProductQualityResponse, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*response.ProductQualityResponse), args.Error(1)
}

func (m *ProductQualityMock) FindAllByProductCode(ctx context.Context, productCode string) (*response.ProductQualityWithOwnProductResponse, error) {
	args := m.Called(ctx, productCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.ProductQualityWithOwnProductResponse), args.Error(1)
}

func (m *ProductQualityMock) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}