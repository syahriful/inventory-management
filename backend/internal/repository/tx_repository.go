package repository

import (
	"context"
	"inventory-management/backend/internal/model"
)

type TxTransactionRepositoryContract interface {
	Create(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error)
	Update(ctx context.Context, requestQuantity float64, transaction *model.Transaction) error
	Delete(ctx context.Context, transaction *model.Transaction) error
}
