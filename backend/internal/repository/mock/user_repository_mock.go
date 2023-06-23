package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/model"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (mock *UserRepositoryMock) FindAll(ctx context.Context) ([]*model.User, error) {
	args := mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*model.User), args.Error(1)
}

func (mock *UserRepositoryMock) FindByID(ctx context.Context, id int64) (*model.User, error) {
	args := mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.User), args.Error(1)
}

func (mock *UserRepositoryMock) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	args := mock.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.User), args.Error(1)
}

func (mock *UserRepositoryMock) Create(ctx context.Context, user *model.User) (*model.User, error) {
	args := mock.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.User), args.Error(1)
}

func (mock *UserRepositoryMock) Update(ctx context.Context, user *model.User) (*model.User, error) {
	args := mock.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.User), args.Error(1)
}

func (mock *UserRepositoryMock) Delete(ctx context.Context, id int64) error {
	args := mock.Called(id)
	return args.Error(0)
}
