package model

import (
	"inventory-management/backend/internal/http/response"
)

type ProductQuality struct {
	ID          int64
	ProductCode string
	Quality     string
	Price       int64
	Quantity    float64
	Type        string
	Product     *Product `gorm:"foreignKey:ProductCode;references:Code"`
}

func (p *ProductQuality) ToResponse() *response.ProductQualityResponse {
	return &response.ProductQualityResponse{
		ID:          p.ID,
		ProductCode: p.ProductCode,
		Quality:     p.Quality,
		Price:       p.Price,
		Quantity:    p.Quantity,
		Type:        p.Type,
	}
}

func (p *ProductQuality) ToResponseWithAssociations() *response.ProductQualityResponse {
	return &response.ProductQualityResponse{
		ID:          p.ID,
		ProductCode: p.ProductCode,
		Quality:     p.Quality,
		Price:       p.Price,
		Quantity:    p.Quantity,
		Type:        p.Type,
		Product:     p.Product.ToResponse(),
	}
}
