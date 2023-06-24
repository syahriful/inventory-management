package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) FindAll(ctx context.Context) ([]*response.UserResponse, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*response.UserResponse), args.Error(1)
}

func (m *UserServiceMock) FindByID(ctx context.Context, id int64) (*response.UserResponse, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.UserResponse), args.Error(1)
}

func (m *UserServiceMock) VerifyLogin(ctx context.Context, request *request.LoginUserRequest) (*response.UserLoginResponse, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.UserLoginResponse), args.Error(1)
}

func (m *UserServiceMock) Create(ctx context.Context, request *request.CreateUserRequest) (*response.UserResponse, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.UserResponse), args.Error(1)
}

func (m *UserServiceMock) Update(ctx context.Context, request *request.UpdateUserRequest) (*response.UserResponse, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.UserResponse), args.Error(1)
}

func (m *UserServiceMock) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
