package repository

import (
	"context"
	"gorm.io/gorm"
	"inventory-management/backend/internal/model"
)

type ProductQualityRepository struct {
	DB *gorm.DB
}

func NewProductQualityRepository(db *gorm.DB) ProductQualityRepositoryContract {
	return &ProductQualityRepository{
		DB: db,
	}
}

func (repository *ProductQualityRepository) FindAll(ctx context.Context) ([]*model.ProductQuality, error) {
	var productQualities []*model.ProductQuality
	err := repository.DB.WithContext(ctx).Find(&productQualities).Error
	if err != nil {
		return nil, err
	}

	return productQualities, nil
}

func (repository *ProductQualityRepository) FindAllByProductCode(ctx context.Context, productCode string) ([]*model.ProductQuality, error) {
	var productQualities []*model.ProductQuality
	err := repository.DB.WithContext(ctx).Where("product_code = ?", productCode).Find(&productQualities).Error
	if err != nil {
		return nil, err
	}

	return productQualities, nil
}

func (repository *ProductQualityRepository) FindByID(ctx context.Context, id int64) (*model.ProductQuality, error) {
	var productQuality model.ProductQuality
	err := repository.DB.WithContext(ctx).First(&productQuality, id).Error
	if err != nil {
		return nil, err
	}

	return &productQuality, nil
}

func (repository *ProductQualityRepository) FindByIDWithAssociations(ctx context.Context, id int64) (*model.ProductQuality, error) {
	var productQuality model.ProductQuality
	err := repository.DB.WithContext(ctx).Preload("Product").First(&productQuality, id).Error
	if err != nil {
		return nil, err
	}

	return &productQuality, nil
}

func (repository *ProductQualityRepository) Delete(ctx context.Context, id int64) error {
	var productQuality model.ProductQuality
	err := repository.DB.WithContext(ctx).Delete(&productQuality, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *ProductQualityRepository) IncreaseStock(ctx context.Context, id int64, quantity float64, tx *gorm.DB) error {
	var productQuality model.ProductQuality
	err := tx.WithContext(ctx).First(&productQuality, id).Error
	if err != nil {
		return err
	}

	productQuality.Quantity += quantity
	err = tx.WithContext(ctx).Select("quantity").Updates(&productQuality).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *ProductQualityRepository) DecreaseStock(ctx context.Context, id int64, quantity float64, tx *gorm.DB) error {
	var productQuality model.ProductQuality
	err := tx.WithContext(ctx).First(&productQuality, id).Error
	if err != nil {
		return err
	}

	productQuality.Quantity -= quantity
	err = tx.WithContext(ctx).Select("quantity").Updates(&productQuality).Error
	if err != nil {
		return err
	}

	return nil
}
