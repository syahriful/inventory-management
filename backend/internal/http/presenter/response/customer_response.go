package response

type CustomerResponse struct {
	ID           int64                  `json:"id"`
	Code         string                 `json:"code"`
	Name         string                 `json:"name"`
	CreatedAt    string                 `json:"created_at"`
	UpdatedAt    string                 `json:"updated_at"`
	Transactions []*TransactionResponse `json:"transactions,omitempty"`
}
