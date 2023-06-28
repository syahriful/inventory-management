package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/model"
	"inventory-management/backend/util"
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

func (repository *TxTransactionRepository) Create(ctx context.Context, request *request.CreateTransactionRequest) (*model.Transaction, error) {
	var createdTransaction *model.Transaction
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		productQuality, err := repository.ProductQualityRepository.FindByIDWithAssociations(ctx, request.ProductQualityID, tx)
		if err != nil {
			return err
		}

		quantity, err := util.CalculateUnitOfMass(productQuality.Product.UnitMassAcronym, request.UnitMassAcronym, request.Quantity)
		if err != nil {
			return err
		}

		var transactionRequest model.Transaction
		transactionRequest.ProductQualityID = request.ProductQualityID
		transactionRequest.SupplierCode = request.SupplierCode
		transactionRequest.CustomerCode = request.CustomerCode
		transactionRequest.Description = request.Description
		transactionRequest.Quantity = request.Quantity
		transactionRequest.Type = request.Type
		transactionRequest.UnitMassAcronym = request.UnitMassAcronym

		trx, err := repository.TransactionRepository.Create(ctx, &transactionRequest, tx)
		if err != nil {
			return err
		}

		if transactionRequest.Type == "IN" {
			err = repository.ProductQualityRepository.IncreaseStock(ctx, transactionRequest.ProductQualityID, quantity, tx)
			if err != nil {
				return err
			}
		}

		if transactionRequest.Type == "OUT" {
			err = repository.ProductQualityRepository.DecreaseStock(ctx, transactionRequest.ProductQualityID, quantity, tx)
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

func (repository *TxTransactionRepository) Update(ctx context.Context, request *request.UpdateTransactionRequest) (*model.Transaction, error) {
	var updatedTransaction *model.Transaction
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		transaction, err := repository.TransactionRepository.FindByCodeWithAssociations(ctx, request.Code, tx)
		if err != nil {
			return err
		}

		if transaction.Type == "TRANSFER" {
			return errors.New(response.ErrorUpdateTransactionTypeTransfer)
		}

		if request.CustomerCode != nil {
			transaction.CustomerCode = request.CustomerCode
		}

		if request.SupplierCode != nil {
			transaction.SupplierCode = request.SupplierCode
		}

		transactionQuantity, err := util.CalculateUnitOfMass(transaction.ProductQuality.Product.UnitMassAcronym, transaction.UnitMassAcronym, transaction.Quantity)
		if err != nil {
			return err
		}

		transaction.Description = request.Description
		transaction.UnitMassAcronym = request.UnitMassAcronym
		transaction.Quantity = transactionQuantity

		increaseStock, err := util.CalculateUnitOfMass(transaction.ProductQuality.Product.UnitMassAcronym, request.UnitMassAcronym, request.Quantity)
		if err != nil {
			return err
		}

		// decrease by current quantity
		err = repository.ProductQualityRepository.DecreaseStock(ctx, transaction.ProductQualityID, transaction.Quantity, tx)
		if err != nil {
			return err
		}

		transaction.Quantity = request.Quantity
		_, err = repository.TransactionRepository.Update(ctx, transaction, tx)
		if err != nil {
			return err
		}

		// increase by request quantity
		err = repository.ProductQualityRepository.IncreaseStock(ctx, transaction.ProductQualityID, increaseStock, tx)
		if err != nil {
			return err
		}

		updatedTransaction = transaction

		return nil
	})
	if err != nil {
		return nil, err
	}

	return updatedTransaction, nil
}
func (repository *TxTransactionRepository) TransferStock(ctx context.Context, request *request.TransferStockTransactionRequest) (*model.Transaction, error) {
	var transaction *model.Transaction
	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		fromQuality, err := repository.ProductQualityRepository.FindByIDWithAssociations(ctx, request.ProductQualityID, tx)
		if err != nil {
			return err
		}

		toQuality, err := repository.ProductQualityRepository.FindByID(ctx, request.ProductQualityIDTransferred, tx)
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

		transaction, err = repository.TransactionRepository.Create(ctx, &transactionRequest, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (repository *TxTransactionRepository) Delete(ctx context.Context, code string) error {
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		transaction, err := repository.TransactionRepository.FindByCodeWithAssociations(ctx, code, tx)
		if err != nil {
			return err
		}

		err = repository.TransactionRepository.Delete(ctx, transaction.Code, tx)
		if err != nil {
			return err
		}

		quantityAfterCalculated, err := util.CalculateUnitOfMass(transaction.ProductQuality.Product.UnitMassAcronym, transaction.UnitMassAcronym, transaction.Quantity)
		if err != nil {
			return err
		}

		if transaction.Type == "TRANSFER" {
			err := repository.ProductQualityRepository.DecreaseStock(ctx, *transaction.ProductQualityIDTransferred, quantityAfterCalculated, tx)
			if err != nil {
				return err
			}

			err = repository.ProductQualityRepository.IncreaseStock(ctx, transaction.ProductQualityID, quantityAfterCalculated, tx)
			if err != nil {
				return err
			}
		}

		if transaction.Type == "IN" {
			err = repository.ProductQualityRepository.DecreaseStock(ctx, transaction.ProductQualityID, quantityAfterCalculated, tx)
			if err != nil {
				return err
			}
		}

		if transaction.Type == "OUT" {
			err = repository.ProductQualityRepository.IncreaseStock(ctx, transaction.ProductQualityID, quantityAfterCalculated, tx)
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
