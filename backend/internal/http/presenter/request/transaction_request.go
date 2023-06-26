package request

type CreateTransactionRequest struct {
	ProductQualityID int64   `json:"product_quality_id" validate:"required,number"`
	SupplierCode     *string `json:"supplier_code" validate:"omitempty,max=100"`
	CustomerCode     *string `json:"customer_code" validate:"omitempty,max=100"`
	Description      *string `json:"description" validate:"omitempty,max=255"`
	Quantity         float64 `json:"quantity" validate:"required,number"`
	Type             string  `json:"type" validate:"required,oneof=IN OUT TRANSFER"`
}

type UpdateTransactionRequest struct {
	Code         string
	CustomerCode *string `json:"customer_code" validate:"omitempty,max=100"`
	SupplierCode *string `json:"supplier_code" validate:"omitempty,max=100"`
	Description  *string `json:"description" validate:"omitempty,max=255"`
	Quantity     float64 `json:"quantity" validate:"required,number"`
}
