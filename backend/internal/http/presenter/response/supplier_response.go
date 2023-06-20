package response

type SupplierResponse struct {
	ID           int64                  `json:"id,omitempty"`
	Code         string                 `json:"code,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Address      string                 `json:"address,omitempty"`
	Phone        string                 `json:"phone,omitempty"`
	CreatedAt    string                 `json:"created_at,omitempty"`
	UpdatedAt    string                 `json:"updated_at,omitempty"`
	Transactions []*TransactionResponse `json:"transactions,omitempty"`
}
