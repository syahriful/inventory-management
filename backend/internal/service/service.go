package service

import (
	"context"
	request2 "inventory-management/backend/internal/http/request"
	response "inventory-management/backend/internal/http/response"
)

type UserServiceContract interface {
	FindAll(ctx context.Context) ([]*response.UserResponse, error)
	FindByID(ctx context.Context, id int64) (*response.UserResponse, error)
	VerifyLogin(ctx context.Context, request *request2.LoginUserRequest) (*response.UserLoginResponse, error)
	Create(ctx context.Context, request *request2.CreateUserRequest) (*response.UserResponse, error)
	Update(ctx context.Context, request *request2.UpdateUserRequest) (*response.UserResponse, error)
	Delete(ctx context.Context, id int64) error
}

type ProductServiceContract interface {
	FindAll(ctx context.Context) ([]*response.ProductResponse, error)
	FindByCode(ctx context.Context, code string) (*response.ProductResponse, error)
	Create(ctx context.Context, request *request2.CreateProductRequest) (*response.ProductResponse, error)
	Update(ctx context.Context, request *request2.UpdateProductRequest) (*response.ProductResponse, error)
	Delete(ctx context.Context, code string) error
}

type ProductQualityServiceContract interface {
	FindAll(ctx context.Context) ([]*response.ProductQualityResponse, error)
	FindAllByProductCode(ctx context.Context, productCode string) (*response.ProductQualityWithOwnProductResponse, error)
	FindByID(ctx context.Context, id int64) (*response.ProductQualityResponse, error)
	Delete(ctx context.Context, id int64) error
}

type SupplierServiceContract interface {
	FindAll(ctx context.Context) ([]*response.SupplierResponse, error)
	FindByCode(ctx context.Context, code string) (*response.SupplierResponse, error)
	Create(ctx context.Context, request *request2.CreateSupplierRequest) (*response.SupplierResponse, error)
	Update(ctx context.Context, request *request2.UpdateSupplierRequest) (*response.SupplierResponse, error)
	Delete(ctx context.Context, code string) error
}

type CustomerServiceContract interface {
	FindAll(ctx context.Context) ([]*response.CustomerResponse, error)
	FindByCode(ctx context.Context, code string) (*response.CustomerResponse, error)
	Create(ctx context.Context, request *request2.CreateCustomerRequest) (*response.CustomerResponse, error)
	Update(ctx context.Context, request *request2.UpdateCustomerRequest) (*response.CustomerResponse, error)
	Delete(ctx context.Context, code string) error
}

type TransactionServiceContract interface {
	FindAll(ctx context.Context) ([]*response.TransactionResponse, error)
	FindAllBySupplierCode(ctx context.Context, supplierCode string) ([]*response.TransactionResponse, error)
	FindAllByCustomerCode(ctx context.Context, customerCode string) ([]*response.TransactionResponse, error)
	FindByCode(ctx context.Context, code string) (*response.TransactionResponse, error)
	Create(ctx context.Context, request *request2.CreateTransactionRequest) (*response.TransactionResponse, error)
	Update(ctx context.Context, request *request2.UpdateTransactionRequest) (*response.TransactionResponse, error)
	Delete(ctx context.Context, code string) error
}
