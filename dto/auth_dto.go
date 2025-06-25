package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	StatusCode int           `json:"status"`
	Message    string        `json:"message"`
	Data       TokenResponse `json:"data"`
}
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterResponse struct {
	StatusCode int      `json:"status"`
	Message    string   `json:"message"`
	Data       AuthData `json:"data"`
}

type LogoutResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type AuthData struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
