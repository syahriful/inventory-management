package model

import (
	"gorm.io/gorm"
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
