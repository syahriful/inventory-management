package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
)

type ProductServiceMock struct {
	mock.Mock
}

func (m *ProductServiceMock) FindAll(ctx context.Context) ([]*response.ProductResponse, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*response.ProductResponse), args.Error(1)
}

func (m *ProductServiceMock) FindByCode(ctx context.Context, code string) (*response.ProductResponse, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.ProductResponse), args.Error(1)
}

func (m *ProductServiceMock) Create(ctx context.Context, request *request.CreateProductRequest) (*response.ProductResponse, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.ProductResponse), args.Error(1)
}

func (m *ProductServiceMock) Update(ctx context.Context, request *request.UpdateProductRequest) (*response.ProductResponse, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.ProductResponse), args.Error(1)
}

func (m *ProductServiceMock) Delete(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	return args.Error(0)
}
