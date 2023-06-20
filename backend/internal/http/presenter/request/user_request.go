package request

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Username string `json:"username" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

type UpdateUserRequest struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"omitempty,min=8,max=255"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
