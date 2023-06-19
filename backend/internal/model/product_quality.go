package model

import "time"

type ProductQuality struct {
	ID          int64
	ProductCode string
	Quality     string
	Price       int64
	Quantity    float64
	Type        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Product     *Product
}
