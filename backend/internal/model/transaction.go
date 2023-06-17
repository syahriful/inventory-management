package model

import "time"

type Transaction struct {
	ID               int64
	Code             string
	ProductQualityID int64
	SupplierCode     *string
	CustomerCode     *string
	Description      *string
	Quantity         float64
	Type             string
	CreatedAt        time.Time
}
