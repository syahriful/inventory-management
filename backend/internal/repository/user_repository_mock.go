package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/model"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) FindAll(ctx context.Context) ([]*model.User, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*model.User), args.Error(1)
}

func (m *UserRepositoryMock) FindByID(ctx context.Context, id int64) (*model.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) Create(ctx context.Context, user *model.User) (*model.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) Update(ctx context.Context, user *model.User) (*model.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) Delete(ctx context.Context, id int64) error {
	args := m.Called(id)
	return args.Error(0)
}
