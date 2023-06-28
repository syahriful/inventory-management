package model

import (
	"gorm.io/gorm"
	response "inventory-management/backend/internal/http/response"
	"inventory-management/backend/util"
	"time"
)

type Transaction struct {
	ID                          int64
	Code                        string
	ProductQualityID            int64
	ProductQuality              *ProductQuality `gorm:"foreignKey:ProductQualityID;references:ID"`
	ProductQualityIDTransferred *int64
	ProductQualityTransferred   *ProductQuality `gorm:"foreignKey:ProductQualityIDTransferred;references:ID"`
	SupplierCode                *string
	Supplier                    *Supplier `gorm:"foreignKey:SupplierCode;references:Code"`
	CustomerCode                *string
	Customer                    *Customer `gorm:"foreignKey:CustomerCode;references:Code"`
	Description                 *string
	Quantity                    float64
	Type                        string
	UnitMassAcronym             string
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
}

func (t *Transaction) setNullable() {
	if t.ProductQualityIDTransferred == nil || *t.ProductQualityIDTransferred == 0 {
		t.ProductQualityIDTransferred = nil
	}

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
		ID:                          t.ID,
		Code:                        t.Code,
		ProductQualityID:            t.ProductQualityID,
		ProductQualityIDTransferred: t.ProductQualityIDTransferred,
		SupplierCode:                t.SupplierCode,
		CustomerCode:                t.CustomerCode,
		Description:                 t.Description,
		Quantity:                    t.Quantity,
		Type:                        t.Type,
		UnitMassAcronym:             t.UnitMassAcronym,
		CreatedAt:                   t.CreatedAt.String(),
		UpdatedAt:                   t.UpdatedAt.String(),
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

	var productQualityTransferredResponse *response.ProductQualityResponse
	if t.ProductQualityTransferred != nil {
		productQualityTransferredResponse = t.ProductQualityTransferred.ToResponse()
	}

	return &response.TransactionResponse{
		ID:                          t.ID,
		Code:                        t.Code,
		ProductQualityID:            t.ProductQualityID,
		ProductQuality:              t.ProductQuality.ToResponseWithAssociations(),
		ProductQualityIDTransferred: t.ProductQualityIDTransferred,
		ProductQualityTransferred:   productQualityTransferredResponse,
		SupplierCode:                t.SupplierCode,
		Supplier:                    supplierResponse,
		CustomerCode:                t.CustomerCode,
		Customer:                    customerResponse,
		Description:                 t.Description,
		Quantity:                    t.Quantity,
		Type:                        t.Type,
		CreatedAt:                   t.CreatedAt.String(),
		UpdatedAt:                   t.UpdatedAt.String(),
	}
}
