package models

// RegisterRequest represents the user registration payload with support for both JSON and form data
type RegisterRequest struct {
	Username string `json:"username" form:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	UsernameorEmail string `json:"username_or_email" form:"username_or_email" validate:"required"`
	Password        string `json:"password" form:"password" validate:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required,min=6"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" binding:"required"`
}

type CheckUsernameRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
}

type TokenResponse struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
