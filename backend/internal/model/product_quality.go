package model

type ProductQuality struct {
	ID          int64
	ProductCode string
	Quality     string
	Price       int64
	Quantity    float64
	Type        string
	Product     *Product `gorm:"foreignKey:ProductCode;references:Code"`
}
