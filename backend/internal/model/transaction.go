package model

import "time"

type Transaction struct {
	ID               int64
	Code             string
	ProductQualityID int64
	ProductQuality   *ProductQuality
	SupplierCode     *string
	Supplier         *Supplier
	CustomerCode     *string
	Customer         *Customer
	Description      *string
	Quantity         float64
	Type             string
	CreatedAt        time.Time
}
