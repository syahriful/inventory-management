package third_party

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/mock"
)

type ElasticsearchMock struct {
	mock.Mock
}

func (service *ElasticsearchMock) CountAll(ctx context.Context, index string, data bytes.Buffer) (int64, error) {
	args := service.Called(ctx, index, data)
	return args.Get(0).(int64), args.Error(1)
}

func (service *ElasticsearchMock) Search(ctx context.Context, index string, data bytes.Buffer, offset int, limit int) (map[string]interface{}, error) {
	args := service.Called(ctx, index, data, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

func (service *ElasticsearchMock) Create(ctx context.Context, index string, request interface{}, id int64) error {
	args := service.Called(ctx, index, request, id)
	return args.Error(0)
}

func (service *ElasticsearchMock) Update(ctx context.Context, index string, request interface{}, id int64) error {
	args := service.Called(ctx, index, request, id)
	return args.Error(0)
}

func (service *ElasticsearchMock) Delete(ctx context.Context, index string, id int64) error {
	args := service.Called(ctx, index, id)
	return args.Error(0)
}
