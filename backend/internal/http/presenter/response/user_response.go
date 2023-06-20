package response

type UserResponse struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Username  string `json:"username,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
