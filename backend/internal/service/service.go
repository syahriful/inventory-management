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
	FindByCode(ctx context.Context, code string) (*response.ProductResponse, error)
	Create(ctx context.Context, request *request.CreateProductRequest) (*response.ProductResponse, error)
	Update(ctx context.Context, request *request.UpdateProductRequest) (*response.ProductResponse, error)
	Delete(ctx context.Context, code string) error
}

type ProductQualityServiceContract interface {
	FindAll(ctx context.Context) ([]*response.ProductQualityResponse, error)
	FindAllByProductCode(ctx context.Context, productCode string) (*response.ProductQualityWithOwnProductResponse, error)
	Delete(ctx context.Context, id int64) error
}

type SupplierServiceContract interface {
	FindAll(ctx context.Context) ([]*response.SupplierResponse, error)
	FindByCode(ctx context.Context, code string) (*response.SupplierResponse, error)
	Create(ctx context.Context, request *request.CreateSupplierRequest) (*response.SupplierResponse, error)
	Update(ctx context.Context, request *request.UpdateSupplierRequest) (*response.SupplierResponse, error)
	Delete(ctx context.Context, code string) error
}

type CustomerServiceContract interface {
	FindAll(ctx context.Context) ([]*response.CustomerResponse, error)
	FindByCode(ctx context.Context, code string) (*response.CustomerResponse, error)
	Create(ctx context.Context, request *request.CreateCustomerRequest) (*response.CustomerResponse, error)
	Update(ctx context.Context, request *request.UpdateCustomerRequest) (*response.CustomerResponse, error)
	Delete(ctx context.Context, code string) error
}
