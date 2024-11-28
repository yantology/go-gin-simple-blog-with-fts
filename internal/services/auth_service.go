package services

import (
	"errors"
	"log"
	"net/http"
	"net/mail"

	"github.com/yantology/go-gin-simple-blog-with-fts/internal/models"
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/utils"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/models/authmodels"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/repositories/authrepository"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/utils/customerror"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo *authrepository.AuthRepository
	jwtUtil  *utils.JWTUtil
}

func NewAuthService(authRepo *authrepository.AuthRepository, jwtUtil *utils.JWTUtil) *AuthService {
	return &AuthService{
		authRepo: authRepo,
		jwtUtil:  jwtUtil,
	}
}

func (s *AuthService) Register(req *models.RegisterRequest) *customerror.CustomError {
	log.Println("Registering user")

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return customerror.NewCustomError(err, "failed to hash password", 500)
	}

	return s.authRepo.CreateUser(req.Username, req.Email, string(hashedPassword))
}

func (s *AuthService) Login(req *models.LoginRequest) (*models.TokenResponse, *customerror.CustomError) {
	//check username or email
	var user *authmodels.User
	var cuserr *customerror.CustomError
	_, err := mail.ParseAddress(req.UsernameorEmail)
	if err != nil {
		user, cuserr = s.authRepo.GetUserByUsername(req.UsernameorEmail)
		err = nil
	} else {
		user, cuserr = s.authRepo.GetUserByEmail(req.UsernameorEmail)
	}

	if cuserr != nil {
		return nil, customerror.NewCustomError(err, "invalid credentials", http.StatusUnauthorized)
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, customerror.NewCustomError(err, "invalid credentials", http.StatusUnauthorized)
	}

	// Generate tokens
	accessToken, refreshToken, err := s.jwtUtil.GenerateTokens(user.ID, user.UpdatedAt)
	if err != nil {
		return nil, customerror.NewCustomError(err, "failed to generate tokens", 500)
	}

	return &models.TokenResponse{
		TokenType:    "Bearer",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) ChangePassword(userID int, req *models.ChangePasswordRequest) *customerror.CustomError {
	// Get user
	user, cuserr := s.authRepo.GetUserByID(userID)
	if cuserr != nil {
		return cuserr
	}

	// Check old password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword))
	if err != nil {
		return customerror.NewCustomError(errors.New("invalid old password"), "invalid old password", http.StatusUnauthorized)
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return customerror.NewCustomError(err, "failed to hash password", 500)
	}

	// Update password
	return s.authRepo.UpdatePassword(userID, string(hashedPassword))
}

func (s *AuthService) RefreshToken(refreshToken string) (*models.TokenResponse, *customerror.CustomError) {
	// Validate refresh token
	token, err := s.jwtUtil.ValidateToken(refreshToken, true)
	if err != nil {
		return nil, customerror.NewCustomError(err, "invalid token", http.StatusUnauthorized)
	}

	// Extract user ID
	userID, err := s.jwtUtil.ExtractUserID(token)
	if err != nil {
		return nil, customerror.NewCustomError(err, "failed to extract user ID", 500)
	}

	// Extract updated at
	updatedAt, err := s.jwtUtil.ExtractUpdatedAt(token)
	if err != nil {
		return nil, customerror.NewCustomError(err, "failed to extract updated at", 500)
	}

	// Get user
	user, cuserr := s.authRepo.GetUserByID(userID)
	if cuserr != nil {
		return nil, cuserr
	}

	// Check if token is still valid
	if updatedAt != user.UpdatedAt {
		return nil, customerror.NewCustomError(nil, "token is no longer valid", http.StatusUnauthorized)
	}

	// Generate new tokens
	accessToken, newRefreshToken, err := s.jwtUtil.GenerateTokens(userID, updatedAt)
	if err != nil {
		return nil, customerror.NewCustomError(err, "failed to generate tokens", 500)
	}

	return &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *AuthService) CheckUsernameExists(username string) *customerror.CustomError {
	statusUsername, cuserr := s.authRepo.CheckUsernameExists(username)
	if cuserr != nil {
		return cuserr
	}

	if statusUsername {
		return customerror.NewCustomError(nil, "username already exists", http.StatusConflict)
	}

	return nil
}
