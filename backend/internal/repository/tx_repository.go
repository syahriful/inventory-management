package repository

import (
	"context"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/model"
)

type TxTransactionRepositoryContract interface {
	Create(ctx context.Context, transaction *model.Transaction, realQuantity float64) (*model.Transaction, error)
	Update(ctx context.Context, increaseStock float64, requestQuantity float64, transaction *model.Transaction) error
	Delete(ctx context.Context, transaction *model.Transaction) error
	TransferStock(ctx context.Context, request *request.TransferStockTransactionRequest) error
}
