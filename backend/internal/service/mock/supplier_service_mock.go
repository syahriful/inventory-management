package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
)

type SupplierServiceMock struct {
	mock.Mock
}

func (mock *SupplierServiceMock) FindAll(ctx context.Context, offset int, limit int) ([]*response.SupplierResponse, error) {
	args := mock.Called(ctx, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*response.SupplierResponse), args.Error(1)
}

func (mock *SupplierServiceMock) CountAll(ctx context.Context) (int64, error) {
	args := mock.Called(ctx)
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}

	return args.Get(0).(int64), args.Error(1)
}

func (mock *SupplierServiceMock) FindByCode(ctx context.Context, code string) (*response.SupplierResponse, error) {
	args := mock.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.SupplierResponse), args.Error(1)
}

func (mock *SupplierServiceMock) Create(ctx context.Context, request *request.CreateSupplierRequest) (*response.SupplierResponse, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.SupplierResponse), args.Error(1)
}

func (mock *SupplierServiceMock) Update(ctx context.Context, request *request.UpdateSupplierRequest) (*response.SupplierResponse, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.SupplierResponse), args.Error(1)
}

func (mock *SupplierServiceMock) Delete(ctx context.Context, code string) error {
	args := mock.Called(ctx, code)
	return args.Error(0)
}
