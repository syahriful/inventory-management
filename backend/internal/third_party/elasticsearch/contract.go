package third_party

import (
	"bytes"
	"context"
)

type ElasticsearchContract interface {
	CountAll(ctx context.Context, index string, data bytes.Buffer) (int64, error)
	Search(ctx context.Context, index string, data bytes.Buffer, offset int, limit int) (map[string]interface{}, error)
	Create(ctx context.Context, index string, request interface{}, id int64) error
	Update(ctx context.Context, index string, request interface{}, id int64) error
	Delete(ctx context.Context, index string, id int64) error
}
