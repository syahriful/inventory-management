package model

import "time"

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
