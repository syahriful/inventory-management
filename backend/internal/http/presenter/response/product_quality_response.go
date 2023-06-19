package response

type ProductQualityResponse struct {
	ID          int64            `json:"id"`
	ProductCode string           `json:"product_code"`
	Quality     string           `json:"quality"`
	Price       int64            `json:"price"`
	Quantity    float64          `json:"quantity"`
	Type        string           `json:"type"`
	Product     *ProductResponse `json:"product,omitempty"`
}
