package model

import (
	"gorm.io/gorm"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/util"
	"time"
)

type Customer struct {
	ID           int64
	Code         string
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Transactions []*Transaction `gorm:"foreignKey:CustomerCode;references:Code"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	c.Code, _ = util.GenerateRandomString(10)

	return nil
}

func (c *Customer) ToResponse() *response.CustomerResponse {
	return &response.CustomerResponse{
		ID:        c.ID,
		Code:      c.Code,
		Name:      c.Name,
		CreatedAt: c.CreatedAt.Local().String(),
		UpdatedAt: c.UpdatedAt.Local().String(),
	}
}

func (c *Customer) ToResponseWithAssociations() *response.CustomerResponse {
	var transactionResponses []*response.TransactionResponse
	for _, transaction := range c.Transactions {
		transactionResponses = append(transactionResponses, transaction.ToResponseWithAssociations())
	}

	return &response.CustomerResponse{
		ID:           c.ID,
		Code:         c.Code,
		Name:         c.Name,
		CreatedAt:    c.CreatedAt.Local().String(),
		UpdatedAt:    c.UpdatedAt.Local().String(),
		Transactions: transactionResponses,
	}
}
