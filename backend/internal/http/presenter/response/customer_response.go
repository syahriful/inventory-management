package response

type CustomerResponse struct {
	ID           int64                  `json:"id,omitempty"`
	Code         string                 `json:"code,omitempty"`
	Name         string                 `json:"name,omitempty"`
	CreatedAt    string                 `json:"created_at,omitempty"`
	UpdatedAt    string                 `json:"updated_at,omitempty"`
	Transactions []*TransactionResponse `json:"transactions,omitempty"`
}
