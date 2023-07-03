package third_party

import (
	"bytes"
	"context"
	"inventory-management/backend/internal/http/response"
)

type UserElasticsearchContract interface {
	CountAll(ctx context.Context, data bytes.Buffer) (int64, error)
	Search(ctx context.Context, data bytes.Buffer, offset int, limit int) (map[string]interface{}, error)
	Create(ctx context.Context, user *response.UserResponse) error
	Update(ctx context.Context, user *response.UserResponse) error
	Delete(ctx context.Context, id int64) error
}
