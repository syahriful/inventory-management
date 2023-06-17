package model

import "time"

type Supplier struct {
	ID        int64
	Code      string
	Name      string
	Address   string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
