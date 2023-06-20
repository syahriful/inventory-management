package model

import (
	"gorm.io/gorm"
	"inventory-management/backend/util"
	"time"
)

type Supplier struct {
	ID           int64
	Code         string
	Name         string
	Address      string
	Phone        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Transactions []*Transaction `gorm:"foreignKey:SupplierCode;references:Code"`
}

func (s *Supplier) BeforeCreate(tx *gorm.DB) error {
	s.Code, _ = util.GenerateRandomString(10)

	return nil
}
