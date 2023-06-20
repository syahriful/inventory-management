package repository

import (
	"context"
	"inventory-management/backend/internal/model"
)

type UserRepositoryContract interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	FindByID(ctx context.Context, id int64) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id int64) error
}

type ProductRepositoryContract interface {
	FindAll(ctx context.Context) ([]*model.Product, error)
	FindByCode(ctx context.Context, code string) (*model.Product, error)
	Create(ctx context.Context, product *model.Product) (*model.Product, error)
	Update(ctx context.Context, product *model.Product) (*model.Product, error)
	Delete(ctx context.Context, code string) error
}

type ProductQualityRepositoryContract interface {
	FindAll(ctx context.Context) ([]*model.ProductQuality, error)
	FindAllByProductCode(ctx context.Context, productCode string) ([]*model.ProductQuality, error)
	FindByID(ctx context.Context, id int64) (*model.ProductQuality, error)
	Delete(ctx context.Context, id int64) error
}

type SupplierRepositoryContract interface {
	FindAll(ctx context.Context) ([]*model.Supplier, error)
	FindByCode(ctx context.Context, code string) (*model.Supplier, error)
	Create(ctx context.Context, supplier *model.Supplier) (*model.Supplier, error)
	Update(ctx context.Context, supplier *model.Supplier) (*model.Supplier, error)
	Delete(ctx context.Context, code string) error
}

type CustomerRepositoryContract interface {
	FindAll(ctx context.Context) ([]*model.Customer, error)
	FindByCode(ctx context.Context, code string) (*model.Customer, error)
	Create(ctx context.Context, customer *model.Customer) (*model.Customer, error)
	Update(ctx context.Context, customer *model.Customer) (*model.Customer, error)
	Delete(ctx context.Context, code string) error
}
