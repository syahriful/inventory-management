package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
)

type CustomerServiceMock struct {
	mock.Mock
}

func (m *CustomerServiceMock) FindAll(ctx context.Context) ([]*response.CustomerResponse, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*response.CustomerResponse), args.Error(1)
}

func (m *CustomerServiceMock) FindByCode(ctx context.Context, code string) (*response.CustomerResponse, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.CustomerResponse), args.Error(1)
}

func (m *CustomerServiceMock) Create(ctx context.Context, request *request.CreateCustomerRequest) (*response.CustomerResponse, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.CustomerResponse), args.Error(1)
}

func (m *CustomerServiceMock) Update(ctx context.Context, request *request.UpdateCustomerRequest) (*response.CustomerResponse, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.CustomerResponse), args.Error(1)
}

func (m *CustomerServiceMock) Delete(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	return args.Error(0)
}
