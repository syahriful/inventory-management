package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/model"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func (mock *CustomerRepositoryMock) FindAll(ctx context.Context) ([]*model.Customer, error) {
	args := mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*model.Customer), args.Error(1)
}

func (mock *CustomerRepositoryMock) FindByCodeWithAssociations(ctx context.Context, code string) (*model.Customer, error) {
	args := mock.Called(code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Customer), args.Error(1)
}

func (mock *CustomerRepositoryMock) FindByCode(ctx context.Context, code string) (*model.Customer, error) {
	args := mock.Called(code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Customer), args.Error(1)
}

func (mock *CustomerRepositoryMock) Create(ctx context.Context, customer *model.Customer) (*model.Customer, error) {
	args := mock.Called(customer)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Customer), args.Error(1)
}

func (mock *CustomerRepositoryMock) Update(ctx context.Context, customer *model.Customer) (*model.Customer, error) {
	args := mock.Called(customer)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Customer), args.Error(1)
}

func (mock *CustomerRepositoryMock) Delete(ctx context.Context, code string) error {
	args := mock.Called(code)
	return args.Error(0)
}
