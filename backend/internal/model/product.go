package model

import (
	"gorm.io/gorm"
	response "inventory-management/backend/internal/http/response"
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

func (p *Product) ToResponse() *response.ProductResponse {
	return &response.ProductResponse{
		ID:                  p.ID,
		Code:                p.Code,
		Name:                p.Name,
		UnitMassAcronym:     p.UnitMassAcronym,
		UnitMassDescription: p.UnitMassDescription,
		CreatedAt:           p.CreatedAt.String(),
		UpdatedAt:           p.UpdatedAt.String(),
	}
}

func (p *Product) ToResponseWithAssociations() *response.ProductResponse {
	var productQualities []*response.ProductQualityResponse
	for _, productQuality := range p.ProductQualities {
		productQualities = append(productQualities, &response.ProductQualityResponse{
			ID:          productQuality.ID,
			ProductCode: productQuality.ProductCode,
			Quality:     productQuality.Quality,
			Price:       productQuality.Price,
			Quantity:    productQuality.Quantity,
			Type:        productQuality.Type,
		})
	}

	return &response.ProductResponse{
		ID:                  p.ID,
		Code:                p.Code,
		Name:                p.Name,
		UnitMassAcronym:     p.UnitMassAcronym,
		UnitMassDescription: p.UnitMassDescription,
		CreatedAt:           p.CreatedAt.String(),
		UpdatedAt:           p.UpdatedAt.String(),
		ProductQualities:    productQualities,
	}
}
