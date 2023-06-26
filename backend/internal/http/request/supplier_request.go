package request

type CreateSupplierRequest struct {
	Name    string `json:"name" validate:"required,min=3,max=100"`
	Address string `json:"address" validate:"required,max=255"`
	Phone   string `json:"phone" validate:"required,min=10,max=15,number"`
}

type UpdateSupplierRequest struct {
	Code    string
	Name    string `json:"name" validate:"required,min=3,max=100"`
	Address string `json:"address" validate:"required,max=255"`
	Phone   string `json:"phone" validate:"required,min=10,max=15,number"`
}
