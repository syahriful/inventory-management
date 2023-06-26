package model

import (
	"gorm.io/gorm"
	response "inventory-management/backend/internal/http/response"
	"inventory-management/backend/util"
	"time"
)

type Transaction struct {
	ID               int64
	Code             string
	ProductQualityID int64
	ProductQuality   *ProductQuality
	SupplierCode     *string
	Supplier         *Supplier `gorm:"foreignKey:SupplierCode;references:Code"`
	CustomerCode     *string
	Customer         *Customer `gorm:"foreignKey:CustomerCode;references:Code"`
	Description      *string
	Quantity         float64
	Type             string
	CreatedAt        time.Time
}

func (t *Transaction) setNullable() {
	if t.SupplierCode == nil || *t.SupplierCode == "" {
		t.SupplierCode = nil
	}

	if t.CustomerCode == nil || *t.CustomerCode == "" {
		t.CustomerCode = nil
	}

	if t.Description == nil || *t.Description == "" {
		t.Description = nil
	}
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	t.Code, _ = util.GenerateRandomString(10)
	t.setNullable()

	return nil
}

func (t *Transaction) BeforeUpdate(tx *gorm.DB) error {
	t.setNullable()

	return nil
}

func (t *Transaction) ToResponse() *response.TransactionResponse {
	return &response.TransactionResponse{
		ID:               t.ID,
		Code:             t.Code,
		ProductQualityID: t.ProductQualityID,
		SupplierCode:     t.SupplierCode,
		CustomerCode:     t.CustomerCode,
		Description:      t.Description,
		Quantity:         t.Quantity,
		Type:             t.Type,
		CreatedAt:        t.CreatedAt.String(),
	}
}

func (t *Transaction) ToResponseWithAssociations() *response.TransactionResponse {
	var supplierResponse *response.SupplierResponse
	if t.Supplier != nil {
		supplierResponse = t.Supplier.ToResponse()
	}

	var customerResponse *response.CustomerResponse
	if t.Customer != nil {
		customerResponse = t.Customer.ToResponse()
	}

	var productQualityResponse *response.ProductQualityResponse
	productQualityResponse = t.ProductQuality.ToResponseWithAssociations()

	return &response.TransactionResponse{
		ID:               t.ID,
		Code:             t.Code,
		ProductQualityID: t.ProductQualityID,
		ProductQuality:   productQualityResponse,
		SupplierCode:     t.SupplierCode,
		Supplier:         supplierResponse,
		CustomerCode:     t.CustomerCode,
		Customer:         customerResponse,
		Description:      t.Description,
		Quantity:         t.Quantity,
		Type:             t.Type,
		CreatedAt:        t.CreatedAt.String(),
	}
}
