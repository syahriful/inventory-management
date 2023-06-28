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

func (repository *TransactionRepository) FindAll(ctx context.Context) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	err := repository.DB.WithContext(ctx).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (repository *TransactionRepository) FindAllBySupplierCode(ctx context.Context, supplierCode string) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	err := repository.DB.WithContext(ctx).Where("supplier_code = ?", supplierCode).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (repository *TransactionRepository) FindAllByCustomerCode(ctx context.Context, customerCode string) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	err := repository.DB.WithContext(ctx).Where("customer_code = ?", customerCode).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (repository *TransactionRepository) FindByCodeWithAssociations(ctx context.Context, code string) (*model.Transaction, error) {
	var transaction model.Transaction
	err := repository.DB.WithContext(ctx).Preload(clause.Associations).Preload("ProductQuality.Product").Where("code = ?", code).First(&transaction).Error
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (repository *TransactionRepository) Create(ctx context.Context, transaction *model.Transaction, tx *gorm.DB) (*model.Transaction, error) {
	err := tx.WithContext(ctx).Omit(clause.Associations).Create(&transaction).Error
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (repository *TransactionRepository) Update(ctx context.Context, transaction *model.Transaction, tx *gorm.DB) (*model.Transaction, error) {
	err := tx.WithContext(ctx).Omit(clause.Associations).Where("code = ?", transaction.Code).Updates(&transaction).Error
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (repository *TransactionRepository) Delete(ctx context.Context, code string, tx *gorm.DB) error {
	var transaction model.Transaction
	err := tx.WithContext(ctx).Where("code = ?", code).Delete(&transaction).Error
	if err != nil {
		return err
	}

	return nil
}
