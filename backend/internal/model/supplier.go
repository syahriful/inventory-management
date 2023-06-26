package model

import (
	"gorm.io/gorm"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/util"
	"time"
)

type Supplier struct {
	ID           int64
	Code         string
	Name         string
	Address      string
	Phone        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Transactions []*Transaction `gorm:"foreignKey:SupplierCode;references:Code"`
}

func (s *Supplier) BeforeCreate(tx *gorm.DB) error {
	s.Code, _ = util.GenerateRandomString(10)

	return nil
}

func (s *Supplier) ToResponse() *response.SupplierResponse {
	return &response.SupplierResponse{
		ID:        s.ID,
		Code:      s.Code,
		Name:      s.Name,
		Address:   s.Address,
		Phone:     s.Phone,
		CreatedAt: s.CreatedAt.String(),
		UpdatedAt: s.UpdatedAt.String(),
	}
}

func (s *Supplier) ToResponseWithAssociations() *response.SupplierResponse {
	var transactionResponses []*response.TransactionResponse
	for _, transaction := range s.Transactions {
		transactionResponses = append(transactionResponses, transaction.ToResponseWithAssociations())
	}

	return &response.SupplierResponse{
		ID:           s.ID,
		Code:         s.Code,
		Name:         s.Name,
		Address:      s.Address,
		Phone:        s.Phone,
		CreatedAt:    s.CreatedAt.String(),
		UpdatedAt:    s.UpdatedAt.String(),
		Transactions: transactionResponses,
	}
}
