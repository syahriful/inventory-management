package request

type CreateCustomerRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type UpdateCustomerRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}
