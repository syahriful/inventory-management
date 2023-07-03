package third_party

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/model"
)

type UserElasticsearchMock struct {
	mock.Mock
}

func (service *UserElasticsearchMock) CountAll(ctx context.Context, data bytes.Buffer) (int64, error) {
	args := service.Called(ctx, data)
	return args.Get(0).(int64), args.Error(1)
}

func (service *UserElasticsearchMock) Search(ctx context.Context, data bytes.Buffer, offset int, limit int) (map[string]interface{}, error) {
	args := service.Called(ctx, data, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

func (service *UserElasticsearchMock) Create(ctx context.Context, user *model.User) error {
	args := service.Called(ctx, user)
	return args.Error(0)
}

func (service *UserElasticsearchMock) Delete(ctx context.Context, id int64) error {
	args := service.Called(ctx, id)
	return args.Error(0)
}
