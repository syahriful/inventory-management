package service

import (
	"bytes"
	"context"
	request "inventory-management/backend/internal/http/request"
	response "inventory-management/backend/internal/http/response"
)

type (
	UserServiceContract interface {
		Search(ctx context.Context, data bytes.Buffer, offset int, limit int, totalRecord chan<- int64) (map[string]interface{}, error)
		FindAll(ctx context.Context, offset int, limit int) ([]*response.UserResponse, error)
		CountAll(ctx context.Context) (int64, error)
		FindByID(ctx context.Context, id int64) (*response.UserResponse, error)
		VerifyLogin(ctx context.Context, request *request.LoginUserRequest) (*response.UserLoginResponse, error)
		Create(ctx context.Context, request *request.CreateUserRequest) (*response.UserResponse, error)
		Update(ctx context.Context, request *request.UpdateUserRequest) (*response.UserResponse, error)
		Delete(ctx context.Context, id int64) error
	}
	ProductServiceContract interface {
		FindAll(ctx context.Context, offset int, limit int) ([]*response.ProductResponse, error)
		CountAll(ctx context.Context) (int64, error)
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
		FindAll(ctx context.Context, offset int, limit int) ([]*response.SupplierResponse, error)
		CountAll(ctx context.Context) (int64, error)
		FindByCode(ctx context.Context, code string) (*response.SupplierResponse, error)
		Create(ctx context.Context, request *request.CreateSupplierRequest) (*response.SupplierResponse, error)
		Update(ctx context.Context, request *request.UpdateSupplierRequest) (*response.SupplierResponse, error)
		Delete(ctx context.Context, code string) error
	}
	CustomerServiceContract interface {
		FindAll(ctx context.Context, offset int, limit int) ([]*response.CustomerResponse, error)
		CountAll(ctx context.Context) (int64, error)
		FindByCode(ctx context.Context, code string) (*response.CustomerResponse, error)
		Create(ctx context.Context, request *request.CreateCustomerRequest) (*response.CustomerResponse, error)
		Update(ctx context.Context, request *request.UpdateCustomerRequest) (*response.CustomerResponse, error)
		Delete(ctx context.Context, code string) error
	}
	TransactionServiceContract interface {
		FindAll(ctx context.Context, offset int, limit int) ([]*response.TransactionResponse, error)
		CountAll(ctx context.Context) (int64, error)
		FindAllBySupplierCode(ctx context.Context, supplierCode string) ([]*response.TransactionResponse, error)
		FindAllByCustomerCode(ctx context.Context, customerCode string) ([]*response.TransactionResponse, error)
		FindByCode(ctx context.Context, code string) (*response.TransactionResponse, error)
		Create(ctx context.Context, request *request.CreateTransactionRequest) (*response.TransactionResponse, error)
		Update(ctx context.Context, request *request.UpdateTransactionRequest) (*response.TransactionResponse, error)
		TransferStock(ctx context.Context, request *request.TransferStockTransactionRequest) (*response.TransactionResponse, error)
		Delete(ctx context.Context, code string) error
	}
)
