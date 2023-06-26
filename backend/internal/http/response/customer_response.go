package response

type CustomerResponse struct {
	ID           int64                  `json:"id"`
	Code         string                 `json:"code"`
	Name         string                 `json:"name"`
	CreatedAt    string                 `json:"created_at,omitempty"`
	UpdatedAt    string                 `json:"updated_at,omitempty"`
	Transactions []*TransactionResponse `json:"transactions,omitempty"`
}
