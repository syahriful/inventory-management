package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
)

type CustomerServiceMock struct {
	mock.Mock
}

func (mock *CustomerServiceMock) FindAll(ctx context.Context, offset int, limit int) ([]*response.CustomerResponse, error) {
	args := mock.Called(ctx, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*response.CustomerResponse), args.Error(1)
}

func (mock *CustomerServiceMock) CountAll(ctx context.Context) (int64, error) {
	args := mock.Called(ctx)
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}

	return args.Get(0).(int64), args.Error(1)
}

func (mock *CustomerServiceMock) FindByCode(ctx context.Context, code string) (*response.CustomerResponse, error) {
	args := mock.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.CustomerResponse), args.Error(1)
}

func (mock *CustomerServiceMock) Create(ctx context.Context, request *request.CreateCustomerRequest) (*response.CustomerResponse, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.CustomerResponse), args.Error(1)
}

func (mock *CustomerServiceMock) Update(ctx context.Context, request *request.UpdateCustomerRequest) (*response.CustomerResponse, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.CustomerResponse), args.Error(1)
}

func (mock *CustomerServiceMock) Delete(ctx context.Context, code string) error {
	args := mock.Called(ctx, code)
	return args.Error(0)
}
