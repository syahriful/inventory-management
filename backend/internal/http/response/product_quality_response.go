package response

type ProductQualityResponse struct {
	ID          int64            `json:"id"`
	ProductCode string           `json:"product_code,omitempty"`
	Quality     string           `json:"quality"`
	Price       int64            `json:"price"`
	Quantity    float64          `json:"quantity"`
	Type        string           `json:"type,omitempty"`
	Product     *ProductResponse `json:"product,omitempty"`
}

type ProductQualityWithOwnProductResponse struct {
	Product          *ProductResponse          `json:"product"`
	ProductQualities []*ProductQualityResponse `json:"product_qualities"`
}
