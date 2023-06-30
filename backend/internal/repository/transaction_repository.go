package repository

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"inventory-management/backend/internal/model"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepositoryContract {
	return &TransactionRepository{
		DB: db,
	}
}

func (repository *TransactionRepository) FindAll(ctx context.Context, offset int, limit int, tx *gorm.DB) ([]*model.Transaction, error) {
	if tx != nil {
		repository.DB = tx
	}

	var transactions []*model.Transaction
	err := repository.DB.WithContext(ctx).Offset(offset).Limit(limit).Order("created_at DESC").Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (repository *TransactionRepository) CountAll(ctx context.Context, tx *gorm.DB) (int64, error) {
	if tx != nil {
		repository.DB = tx
	}

	var count int64
	err := repository.DB.WithContext(ctx).Model(&model.Transaction{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (repository *TransactionRepository) FindAllBySupplierCode(ctx context.Context, supplierCode string, tx *gorm.DB) ([]*model.Transaction, error) {
	if tx != nil {
		repository.DB = tx
	}

	var transactions []*model.Transaction
	err := repository.DB.WithContext(ctx).Where("supplier_code = ?", supplierCode).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (repository *TransactionRepository) FindAllByCustomerCode(ctx context.Context, customerCode string, tx *gorm.DB) ([]*model.Transaction, error) {
	if tx != nil {
		repository.DB = tx
	}

	var transactions []*model.Transaction
	err := repository.DB.WithContext(ctx).Where("customer_code = ?", customerCode).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (repository *TransactionRepository) FindByCodeWithAssociations(ctx context.Context, code string, tx *gorm.DB) (*model.Transaction, error) {
	if tx != nil {
		repository.DB = tx
	}

	var transaction model.Transaction
	err := repository.DB.WithContext(ctx).Preload(clause.Associations).Preload("ProductQuality.Product").Preload("ProductQualityTransferred.Product").Where("code = ?", code).First(&transaction).Error
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (repository *TransactionRepository) FindByCode(ctx context.Context, code string, tx *gorm.DB) (*model.Transaction, error) {
	if tx != nil {
		repository.DB = tx
	}

	var transaction model.Transaction
	err := repository.DB.WithContext(ctx).Where("code = ?", code).First(&transaction).Error
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (repository *TransactionRepository) Create(ctx context.Context, transaction *model.Transaction, tx *gorm.DB) (*model.Transaction, error) {
	if tx != nil {
		repository.DB = tx
	}

	err := repository.DB.WithContext(ctx).Omit(clause.Associations).Create(&transaction).Error
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (repository *TransactionRepository) Update(ctx context.Context, transaction *model.Transaction, tx *gorm.DB) (*model.Transaction, error) {
	if tx != nil {
		repository.DB = tx
	}

	err := repository.DB.WithContext(ctx).Omit(clause.Associations).Where("code = ?", transaction.Code).Updates(&transaction).Error
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (repository *TransactionRepository) Delete(ctx context.Context, code string, tx *gorm.DB) error {
	if tx != nil {
		repository.DB = tx
	}

	var transaction model.Transaction
	err := repository.DB.WithContext(ctx).Where("code = ?", code).Delete(&transaction).Error
	if err != nil {
		return err
	}

	return nil
}
