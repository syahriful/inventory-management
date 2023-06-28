package service

import (
	"context"
	request "inventory-management/backend/internal/http/request"
	response "inventory-management/backend/internal/http/response"
)

type (
	UserServiceContract interface {
		FindAll(ctx context.Context) ([]*response.UserResponse, error)
		FindByID(ctx context.Context, id int64) (*response.UserResponse, error)
		VerifyLogin(ctx context.Context, request *request.LoginUserRequest) (*response.UserLoginResponse, error)
		Create(ctx context.Context, request *request.CreateUserRequest) (*response.UserResponse, error)
		Update(ctx context.Context, request *request.UpdateUserRequest) (*response.UserResponse, error)
		Delete(ctx context.Context, id int64) error
	}
	ProductServiceContract interface {
		FindAll(ctx context.Context) ([]*response.ProductResponse, error)
		FindByCode(ctx context.Context, code string) (*response.ProductResponse, error)
		Create(ctx context.Context, request *request.CreateProductRequest) (*response.ProductResponse, error)
		Update(ctx context.Context, request *request.UpdateProductRequest) (*response.ProductResponse, error)
		Delete(ctx context.Context, code string) error
	}
	ProductQualityServiceContract interface {
		FindAll(ctx context.Context) ([]*response.ProductQualityResponse, error)
		FindAllByProductCode(ctx context.Context, productCode string) (*response.ProductQualityWithOwnProductResponse, error)
		FindByID(ctx context.Context, id int64) (*response.ProductQualityResponse, error)
		Delete(ctx context.Context, id int64) error
	}
	SupplierServiceContract interface {
		FindAll(ctx context.Context) ([]*response.SupplierResponse, error)
		FindByCode(ctx context.Context, code string) (*response.SupplierResponse, error)
		Create(ctx context.Context, request *request.CreateSupplierRequest) (*response.SupplierResponse, error)
		Update(ctx context.Context, request *request.UpdateSupplierRequest) (*response.SupplierResponse, error)
		Delete(ctx context.Context, code string) error
	}
	CustomerServiceContract interface {
		FindAll(ctx context.Context) ([]*response.CustomerResponse, error)
		FindByCode(ctx context.Context, code string) (*response.CustomerResponse, error)
		Create(ctx context.Context, request *request.CreateCustomerRequest) (*response.CustomerResponse, error)
		Update(ctx context.Context, request *request.UpdateCustomerRequest) (*response.CustomerResponse, error)
		Delete(ctx context.Context, code string) error
	}
	TransactionServiceContract interface {
		FindAll(ctx context.Context) ([]*response.TransactionResponse, error)
		FindAllBySupplierCode(ctx context.Context, supplierCode string) ([]*response.TransactionResponse, error)
		FindAllByCustomerCode(ctx context.Context, customerCode string) ([]*response.TransactionResponse, error)
		FindByCode(ctx context.Context, code string) (*response.TransactionResponse, error)
		Create(ctx context.Context, request *request.CreateTransactionRequest) (*response.TransactionResponse, error)
		Update(ctx context.Context, request *request.UpdateTransactionRequest) (*response.TransactionResponse, error)
		TransferStock(ctx context.Context, request *request.TransferStockTransactionRequest) (*response.TransactionResponse, error)
		Delete(ctx context.Context, code string) error
	}
)
