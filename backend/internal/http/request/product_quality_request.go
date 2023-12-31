package request

type CreateProductQualityRequest struct {
	Quality  string  `json:"quality" validate:"required,max=100"`
	Price    int64   `json:"price" validate:"required,number"`
	Quantity float64 `json:"quantity" validate:"omitempty,number"`
	Type     string  `json:"type" validate:"required,max=20"`
}

type UpdateProductQualityRequest struct {
	ID       int64
	Quality  string  `json:"quality" validate:"required,max=100"`
	Price    int64   `json:"price" validate:"required,number"`
	Quantity float64 `json:"quantity" validate:"omitempty,number"`
	Type     string  `json:"type" validate:"required,max=20"`
}
