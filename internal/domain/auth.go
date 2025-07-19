package domain

type LoginRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=20"`
	Password string `json:"password" binding:"required"`
}

// Response struct for login
type LoginResponse struct {
	User  User          `json:"user"`
	Token TokenResponse `json:"info"`
}

type TokenResponse struct {
	Token        string `json:"token"`
	ExpireAt     string `json:"expire_at"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type"`
	IssuedAt     string `json:"issued_at"`
}

// Request struct for registration
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" validate:"required,alphanum,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	RoleID   *uint  `json:"role_id,omitempty"` // default user
}

// Response after registration
type RegisterResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
