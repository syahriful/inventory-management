package request

type CreateCustomerRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type UpdateCustomerRequest struct {
	Code string
	Name string `json:"name" validate:"required,min=3,max=100"`
}
