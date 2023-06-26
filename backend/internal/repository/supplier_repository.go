package repository

import (
	"context"
	"gorm.io/gorm"
	"inventory-management/backend/internal/model"
)

type SupplierRepository struct {
	DB *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepositoryContract {
	return &SupplierRepository{
		DB: db,
	}
}

func (repository *SupplierRepository) FindAll(ctx context.Context) ([]*model.Supplier, error) {
	var suppliers []*model.Supplier
	err := repository.DB.WithContext(ctx).Find(&suppliers).Error
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (repository *SupplierRepository) FindByCodeWithAssociations(ctx context.Context, code string) (*model.Supplier, error) {
	var supplier model.Supplier
	err := repository.DB.WithContext(ctx).Preload("Transactions").Preload("Transactions.ProductQuality", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id", "product_code", "quality", "quantity", "price")
	}).Preload("Transactions.ProductQuality.Product").Where("code = ?", code).First(&supplier).Error
	if err != nil {
		return nil, err
	}

	return &supplier, nil
}

func (repository *SupplierRepository) FindByCode(ctx context.Context, code string) (*model.Supplier, error) {
	var supplier model.Supplier
	err := repository.DB.WithContext(ctx).Where("code = ?", code).First(&supplier).Error
	if err != nil {
		return nil, err
	}

	return &supplier, nil
}

func (repository *SupplierRepository) Create(ctx context.Context, supplier *model.Supplier) (*model.Supplier, error) {
	err := repository.DB.WithContext(ctx).Create(supplier).Error
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (repository *SupplierRepository) Update(ctx context.Context, supplier *model.Supplier) (*model.Supplier, error) {
	err := repository.DB.WithContext(ctx).Where("code = ?", supplier.Code).Updates(&supplier).Error
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (repository *SupplierRepository) Delete(ctx context.Context, code string) error {
	var supplier model.Supplier
	err := repository.DB.WithContext(ctx).Where("code = ?", code).Delete(&supplier).Error
	if err != nil {
		return err
	}

	return nil
}
