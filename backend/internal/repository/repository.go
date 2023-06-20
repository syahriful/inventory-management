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
