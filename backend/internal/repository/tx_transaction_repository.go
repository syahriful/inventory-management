package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
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

func (repository *TxTransactionRepository) Create(ctx context.Context, transaction *model.Transaction, realQuantity float64) (*model.Transaction, error) {
	var createdTransaction *model.Transaction
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		trx, err := repository.TransactionRepository.Create(ctx, transaction, tx)
		if err != nil {
			return err
		}

		if transaction.Type == "IN" {
			err = repository.ProductQualityRepository.IncreaseStock(ctx, transaction.ProductQualityID, realQuantity, tx)
			if err != nil {
				return err
			}
		}

		if transaction.Type == "OUT" {
			err = repository.ProductQualityRepository.DecreaseStock(ctx, transaction.ProductQualityID, realQuantity, tx)
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

func (repository *TxTransactionRepository) Update(ctx context.Context, increaseStock float64, requestQuantity float64, transaction *model.Transaction) error {
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
		err = repository.ProductQualityRepository.IncreaseStock(ctx, transaction.ProductQualityID, increaseStock, tx)
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

		if transaction.Type == "TRANSFER" {
			err := repository.ProductQualityRepository.DecreaseStock(ctx, *transaction.ProductQualityIDTransferred, transaction.Quantity, tx)
			if err != nil {
				return err
			}

			err = repository.ProductQualityRepository.IncreaseStock(ctx, transaction.ProductQualityID, transaction.Quantity, tx)
			if err != nil {
				return err
			}
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

func (repository *TxTransactionRepository) TransferStock(ctx context.Context, request *request.TransferStockTransactionRequest) error {
	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		fromQuality, err := repository.ProductQualityRepository.FindByIDWithAssociations(ctx, request.ProductQualityID)
		if err != nil {
			return err
		}

		toQuality, err := repository.ProductQualityRepository.FindByID(ctx, request.ProductQualityIDTransferred)
		if err != nil {
			return err
		}

		if fromQuality.ProductCode != toQuality.ProductCode {
			return errors.New(response.ErrorTransferStockDifferentProduct)
		}

		if fromQuality.Quantity < request.Quantity {
			return errors.New(response.ErrorStockNotEnough)
		}

		err = repository.ProductQualityRepository.DecreaseStock(ctx, fromQuality.ID, request.Quantity, tx)
		if err != nil {
			return err
		}

		err = repository.ProductQualityRepository.IncreaseStock(ctx, toQuality.ID, request.Quantity, tx)
		if err != nil {
			return err
		}

		var transactionRequest model.Transaction
		transactionRequest.ProductQualityID = request.ProductQualityID
		transactionRequest.ProductQualityIDTransferred = &request.ProductQualityIDTransferred
		transactionRequest.Description = request.Description
		transactionRequest.Quantity = request.Quantity
		transactionRequest.Type = "TRANSFER"
		transactionRequest.UnitMassAcronym = fromQuality.Product.UnitMassAcronym

		_, err = repository.TransactionRepository.Create(ctx, &transactionRequest, tx)
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
