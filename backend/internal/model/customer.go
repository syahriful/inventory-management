package model

import "time"

type Customer struct {
	ID           int64
	Code         string
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Transactions []*Transaction
}
