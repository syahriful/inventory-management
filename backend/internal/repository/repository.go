package repository

import (
	"context"
	"gorm.io/gorm"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/model"
)

type (
	UserRepositoryContract interface {
		FindAll(ctx context.Context, offset int, limit int) ([]*model.User, error)
		CountAll(ctx context.Context) (int64, error)
		FindByID(ctx context.Context, id int64) (*model.User, error)
		FindByUsername(ctx context.Context, username string) (*model.User, error)
		Create(ctx context.Context, user *model.User) (*model.User, error)
		Update(ctx context.Context, user *model.User) (*model.User, error)
		Delete(ctx context.Context, id int64) error
	}
	ProductRepositoryContract interface {
		FindAll(ctx context.Context, offset int, limit int) ([]*model.Product, error)
		CountAll(ctx context.Context) (int64, error)
		FindByCodeWithAssociations(ctx context.Context, code string) (*model.Product, error)
		Create(ctx context.Context, product *model.Product) (*model.Product, error)
		Update(ctx context.Context, product *model.Product) (*model.Product, error)
		Delete(ctx context.Context, code string) error
	}
	ProductQualityRepositoryContract interface {
		FindAll(ctx context.Context, tx *gorm.DB) ([]*model.ProductQuality, error)
		FindAllByProductCode(ctx context.Context, productCode string, tx *gorm.DB) ([]*model.ProductQuality, error)
		FindByID(ctx context.Context, id int64, tx *gorm.DB) (*model.ProductQuality, error)
		FindByIDWithAssociations(ctx context.Context, id int64, tx *gorm.DB) (*model.ProductQuality, error)
		Delete(ctx context.Context, id int64, tx *gorm.DB) error
		IncreaseStock(ctx context.Context, id int64, quantity float64, tx *gorm.DB) error
		DecreaseStock(ctx context.Context, id int64, quantity float64, tx *gorm.DB) error
	}
	SupplierRepositoryContract interface {
		FindAll(ctx context.Context, offset int, limit int) ([]*model.Supplier, error)
		CountAll(ctx context.Context) (int64, error)
		FindByCodeWithAssociations(ctx context.Context, code string) (*model.Supplier, error)
		FindByCode(ctx context.Context, code string) (*model.Supplier, error)
		Create(ctx context.Context, supplier *model.Supplier) (*model.Supplier, error)
		Update(ctx context.Context, supplier *model.Supplier) (*model.Supplier, error)
		Delete(ctx context.Context, code string) error
	}

	CustomerRepositoryContract interface {
		FindAll(ctx context.Context, offset int, limit int) ([]*model.Customer, error)
		CountAll(ctx context.Context) (int64, error)
		FindByCodeWithAssociations(ctx context.Context, code string) (*model.Customer, error)
		FindByCode(ctx context.Context, code string) (*model.Customer, error)
		Create(ctx context.Context, customer *model.Customer) (*model.Customer, error)
		Update(ctx context.Context, customer *model.Customer) (*model.Customer, error)
		Delete(ctx context.Context, code string) error
	}

	TransactionRepositoryContract interface {
		FindAll(ctx context.Context, offset int, limit int, tx *gorm.DB) ([]*model.Transaction, error)
		CountAll(ctx context.Context, tx *gorm.DB) (int64, error)
		FindAllBySupplierCode(ctx context.Context, supplierCode string, tx *gorm.DB) ([]*model.Transaction, error)
		FindAllByCustomerCode(ctx context.Context, customerCode string, tx *gorm.DB) ([]*model.Transaction, error)
		FindByCodeWithAssociations(ctx context.Context, code string, tx *gorm.DB) (*model.Transaction, error)
		FindByCode(ctx context.Context, code string, tx *gorm.DB) (*model.Transaction, error)
		Create(ctx context.Context, transaction *model.Transaction, tx *gorm.DB) (*model.Transaction, error)
		Update(ctx context.Context, transaction *model.Transaction, tx *gorm.DB) (*model.Transaction, error)
		Delete(ctx context.Context, code string, tx *gorm.DB) error
	}

	TxTransactionRepositoryContract interface {
		Create(ctx context.Context, request *request.CreateTransactionRequest) (*model.Transaction, error)
		Update(ctx context.Context, request *request.UpdateTransactionRequest) (*model.Transaction, error)
		TransferStock(ctx context.Context, request *request.TransferStockTransactionRequest) (*model.Transaction, error)
		Delete(ctx context.Context, code string) error
	}
)
