package service

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
)

type UserServiceMock struct {
	mock.Mock
}

func (mock *UserServiceMock) Search(ctx context.Context, data bytes.Buffer, offset int, limit int, totalRecord chan<- int64) (map[string]interface{}, error) {
	args := mock.Called(ctx, data, offset, limit, totalRecord)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(map[string]interface{}), args.Error(1)
}

func (mock *UserServiceMock) FindAll(ctx context.Context, offset int, limit int) ([]*response.UserResponse, error) {
	args := mock.Called(ctx, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*response.UserResponse), args.Error(1)
}

func (mock *UserServiceMock) CountAll(ctx context.Context) (int64, error) {
	args := mock.Called(ctx)
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}

	return args.Get(0).(int64), args.Error(1)
}

func (mock *UserServiceMock) FindByID(ctx context.Context, id int64) (*response.UserResponse, error) {
	args := mock.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.UserResponse), args.Error(1)
}

func (mock *UserServiceMock) VerifyLogin(ctx context.Context, request *request.LoginUserRequest) (*response.UserLoginResponse, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.UserLoginResponse), args.Error(1)
}

func (mock *UserServiceMock) Create(ctx context.Context, request *request.CreateUserRequest) (*response.UserResponse, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.UserResponse), args.Error(1)
}

func (mock *UserServiceMock) Update(ctx context.Context, request *request.UpdateUserRequest) (*response.UserResponse, error) {
	args := mock.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*response.UserResponse), args.Error(1)
}

func (mock *UserServiceMock) Delete(ctx context.Context, id int64) error {
	args := mock.Called(ctx, id)
	return args.Error(0)
}
