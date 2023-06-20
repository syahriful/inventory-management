package model

import (
	"gorm.io/gorm"
	"inventory-management/backend/util"
	"time"
)

type Customer struct {
	ID           int64
	Code         string
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Transactions []*Transaction `gorm:"foreignKey:CustomerCode;references:Code"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	c.Code, _ = util.GenerateRandomString(10)

	return nil
}
