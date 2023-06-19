package service

import (
	"context"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
)

type UserServiceContract interface {
	FindAll(ctx context.Context) ([]*response.UserResponse, error)
	FindByID(ctx context.Context, id int64) (*response.UserResponse, error)
	Create(ctx context.Context, request *request.CreateUserRequest) (*response.UserResponse, error)
	Update(ctx context.Context, request *request.UpdateUserRequest) (*response.UserResponse, error)
	Delete(ctx context.Context, id int64) error
}

type ProductServiceContract interface {
	FindAll(ctx context.Context) ([]*response.ProductResponse, error)
	FindByID(ctx context.Context, id int64) (*response.ProductResponse, error)
	FindByCode(ctx context.Context, code string) (*response.ProductResponse, error)
	Create(ctx context.Context, request *request.CreateProductRequest) (*response.ProductResponse, error)
	Update(ctx context.Context, request *request.UpdateProductRequest) (*response.ProductResponse, error)
	Delete(ctx context.Context, id int64) error
}
