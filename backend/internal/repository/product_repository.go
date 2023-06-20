package repository

import (
	"context"
	"gorm.io/gorm"
	"inventory-management/backend/internal/model"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryContract {
	return &ProductRepository{
		DB: db,
	}
}

func (repository *ProductRepository) FindAll(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product
	err := repository.DB.WithContext(ctx).Model(&model.Product{}).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repository *ProductRepository) FindByCode(ctx context.Context, code string) (*model.Product, error) {
	var product model.Product
	err := repository.DB.WithContext(ctx).Preload("ProductQualities").Where("code = ?", code).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repository *ProductRepository) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Create(&product).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repository *ProductRepository) Update(ctx context.Context, product *model.Product) (*model.Product, error) {
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Where("code = ?", product.Code).Updates(&product).Error
		if err != nil {
			return err
		}

		for _, pq := range product.ProductQualities {
			pq.ProductCode = product.Code
			err = tx.WithContext(ctx).Save(pq).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repository *ProductRepository) Delete(ctx context.Context, code string) error {
	var product model.Product
	err := repository.DB.WithContext(ctx).Where("code = ?", code).Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}
