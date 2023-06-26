package repository

import (
	"context"
	"gorm.io/gorm"
	"inventory-management/backend/internal/model"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepositoryContract {
	return &CustomerRepository{
		DB: db,
	}
}

func (repository *CustomerRepository) FindAll(ctx context.Context) ([]*model.Customer, error) {
	var customers []*model.Customer
	err := repository.DB.WithContext(ctx).Find(&customers).Error
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (repository *CustomerRepository) FindByCodeWithAssociations(ctx context.Context, code string) (*model.Customer, error) {
	var customer model.Customer
	err := repository.DB.WithContext(ctx).Preload("Transactions").Preload("Transactions.ProductQuality", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id", "product_code", "quality", "price")
	}).Preload("Transactions.ProductQuality.Product").Where("code = ?", code).First(&customer).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (repository *CustomerRepository) FindByCode(ctx context.Context, code string) (*model.Customer, error) {
	var customer model.Customer
	err := repository.DB.WithContext(ctx).Where("code = ?", code).First(&customer).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (repository *CustomerRepository) Create(ctx context.Context, customer *model.Customer) (*model.Customer, error) {
	err := repository.DB.WithContext(ctx).Create(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repository *CustomerRepository) Update(ctx context.Context, customer *model.Customer) (*model.Customer, error) {
	err := repository.DB.WithContext(ctx).Where("code = ?", customer.Code).Updates(&customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repository *CustomerRepository) Delete(ctx context.Context, code string) error {
	var customer model.Customer
	err := repository.DB.WithContext(ctx).Where("code = ?", code).Delete(&customer).Error
	if err != nil {
		return err
	}

	return nil
}
