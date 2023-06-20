package response

type ProductQualityResponse struct {
	ID          int64            `json:"id,omitempty"`
	ProductCode string           `json:"product_code,omitempty"`
	Quality     string           `json:"quality,omitempty"`
	Price       int64            `json:"price,omitempty"`
	Quantity    float64          `json:"quantity,omitempty"`
	Type        string           `json:"type,omitempty"`
	Product     *ProductResponse `json:"product,omitempty"`
}

type ProductQualityWithOwnProductResponse struct {
	Product          *ProductResponse          `json:"product"`
	ProductQualities []*ProductQualityResponse `json:"product_qualities"`
}
