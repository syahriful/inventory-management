package repository

import (
	"context"
	"gorm.io/gorm"
	"inventory-management/backend/internal/model"
)

type TxTransactionRepository struct {
	DB                       *gorm.DB
	TransactionRepository    TransactionRepositoryContract
	ProductQualityRepository ProductQualityRepositoryContract
}

func NewTxRepository(db *gorm.DB, transactionRepository TransactionRepositoryContract, productQualityRepository ProductQualityRepositoryContract) TxTransactionRepositoryContract {
	return &TxTransactionRepository{
		DB:                       db,
		TransactionRepository:    transactionRepository,
		ProductQualityRepository: productQualityRepository,
	}
}

func (repository *TxTransactionRepository) Create(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	var createdTransaction *model.Transaction
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		trx, err := repository.TransactionRepository.Create(ctx, transaction, tx)
		if err != nil {
			return err
		}

		if transaction.Type == "IN" {
			err = repository.ProductQualityRepository.IncreaseStock(ctx, transaction.ProductQualityID, transaction.Quantity, tx)
			if err != nil {
				return err
			}
		}

		if transaction.Type == "OUT" {
			err = repository.ProductQualityRepository.DecreaseStock(ctx, transaction.ProductQualityID, transaction.Quantity, tx)
			if err != nil {
				return err
			}
		}

		createdTransaction = trx

		return nil
	})
	if err != nil {
		return nil, err
	}

	return createdTransaction, nil
}

func (repository *TxTransactionRepository) Update(ctx context.Context, requestQuantity float64, transaction *model.Transaction) error {
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// decrease by current quantity
		err := repository.ProductQualityRepository.DecreaseStock(ctx, transaction.ProductQualityID, transaction.Quantity, tx)
		if err != nil {
			return err
		}

		transaction.Quantity = requestQuantity
		_, err = repository.TransactionRepository.Update(ctx, transaction, tx)
		if err != nil {
			return err
		}

		// increase by request quantity
		err = repository.ProductQualityRepository.IncreaseStock(ctx, transaction.ProductQualityID, requestQuantity, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (repository *TxTransactionRepository) Delete(ctx context.Context, transaction *model.Transaction) error {
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := repository.TransactionRepository.Delete(ctx, transaction.Code, tx)
		if err != nil {
			return err
		}

		if transaction.Type == "IN" {
			err = repository.ProductQualityRepository.DecreaseStock(ctx, transaction.ProductQualityID, transaction.Quantity, tx)
			if err != nil {
				return err
			}
		}

		if transaction.Type == "OUT" {
			err = repository.ProductQualityRepository.IncreaseStock(ctx, transaction.ProductQualityID, transaction.Quantity, tx)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
