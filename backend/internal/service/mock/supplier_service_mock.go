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

func (m *SupplierServiceMock) FindAll(ctx context.Context) ([]*response.SupplierResponse, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*response.SupplierResponse), args.Error(1)
}

func (m *SupplierServiceMock) FindByCode(ctx context.Context, code string) (*response.SupplierResponse, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.SupplierResponse), args.Error(1)
}

func (m *SupplierServiceMock) Create(ctx context.Context, request *request.CreateSupplierRequest) (*response.SupplierResponse, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.SupplierResponse), args.Error(1)
}

func (m *SupplierServiceMock) Update(ctx context.Context, request *request.UpdateSupplierRequest) (*response.SupplierResponse, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.SupplierResponse), args.Error(1)
}

func (m *SupplierServiceMock) Delete(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	return args.Error(0)
}
