package model

import (
	"gorm.io/gorm"
	"inventory-management/backend/util"
	"time"
)

type Product struct {
	ID                  int64
	Code                string
	Name                string
	UnitMassAcronym     string
	UnitMassDescription string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	ProductQualities    []*ProductQuality `gorm:"foreignKey:ProductCode;references:Code"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.Code, _ = util.GenerateRandomString(10)

	return nil
}
