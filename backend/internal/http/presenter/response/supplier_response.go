package response

type SupplierResponse struct {
	ID           int64                  `json:"id"`
	Code         string                 `json:"code"`
	Name         string                 `json:"name"`
	Address      string                 `json:"address"`
	Phone        string                 `json:"phone"`
	CreatedAt    string                 `json:"created_at"`
	UpdatedAt    string                 `json:"updated_at"`
	Transactions []*TransactionResponse `json:"transactions,omitempty"`
}
