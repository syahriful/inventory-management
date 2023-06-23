package response

type TransactionResponse struct {
	ID               int64                   `json:"id,omitempty"`
	Code             string                  `json:"code,omitempty"`
	ProductQualityID int64                   `json:"product_quality_id,omitempty"`
	ProductQuality   *ProductQualityResponse `json:"product_quality,omitempty"`
	SupplierCode     *string                 `json:"supplier_code,omitempty"`
	Supplier         *SupplierResponse       `json:"supplier,omitempty"`
	CustomerCode     *string                 `json:"customer_code,omitempty"`
	Customer         *CustomerResponse       `json:"customer,omitempty"`
	Description      *string                 `json:"description,omitempty"`
	Quantity         float64                 `json:"quantity,omitempty"`
	Type             string                  `json:"type,omitempty"`
	CreatedAt        string                  `json:"created_at,omitempty"`
}