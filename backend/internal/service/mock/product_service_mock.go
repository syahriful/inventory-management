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

func (mock *ProductServiceMock) FindAll(ctx context.Context, offset int, limit int) ([]*response.ProductResponse, error) {
	args := mock.Called(ctx, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*response.ProductResponse), args.Error(1)
}

func (mock *ProductServiceMock) CountAll(ctx context.Context) (int64, error) {
	args := mock.Called(ctx)
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}

	return args.Get(0).(int64), args.Error(1)
}

func (mock *ProductServiceMock) FindByCode(ctx context.Context, code string) (*response.ProductResponse, error) {
	args := mock.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.ProductResponse), args.Error(1)
}

func (mock *ProductServiceMock) Create(ctx context.Context, request *request.CreateProductRequest) (*response.ProductResponse, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.ProductResponse), args.Error(1)
}

func (mock *ProductServiceMock) Update(ctx context.Context, request *request.UpdateProductRequest) (*response.ProductResponse, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.ProductResponse), args.Error(1)
}

func (mock *ProductServiceMock) Delete(ctx context.Context, code string) error {
	args := mock.Called(ctx, code)
	return args.Error(0)
}
