package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/models"
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register registers a new user.
// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Param register body models.RegisterRequest false "Register Request"
// @Param username formData string false "Username"
// @Param email formData string false "Email"
// @Param password formData string false "Password"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
// @Router /register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req *models.RegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewMessage(err.Error()))
		return
	}

	if cuserr := h.authService.Register(req); cuserr != nil {
		c.JSON(cuserr.HTTPCode, models.NewMessage(cuserr.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
}

// Login logs in a user.
// @Summary Login a user
// @Description Login a user
// @Tags auth
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Param login body models.LoginRequest false "Login Request"
// @Param username_or_email formData string false "Username or Email"
// @Param password formData string false "Password"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 500 {object} models.Message
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req *models.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewMessage(err.Error()))
		return
	}

	token, cuserr := h.authService.Login(req)
	if cuserr != nil {
		c.JSON(cuserr.HTTPCode, models.NewMessage(cuserr.Error()))
		return
	}

	c.JSON(http.StatusOK, token)
}

// RefreshToken refreshes a user's token.
// @Summary Refresh a user's token
// @Description Refresh a user's token
// @Tags auth
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Param refreshToken body models.RefreshTokenRequest false "Refresh Token Request"
// @Param refresh_token formData string false "Refresh Token"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 500 {object} models.Message
// @Router /refresh-token [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req *models.RefreshTokenRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewMessage(err.Error()))
		return
	}

	token, cuserr := h.authService.RefreshToken(req.RefreshToken)
	if cuserr != nil {
		c.JSON(cuserr.HTTPCode, models.NewMessage(cuserr.Error()))
		return
	}

	c.JSON(http.StatusOK, token)
}

// ChangePassword changes a user's password.
// @Summary Change a user's password
// @Description Change a user's password
// @Tags auth
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Param changePassword body models.ChangePasswordRequest false "Change Password Request"
// @Param old_password formData string false "Old Password"
// @Param new_password formData string false "New Password"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 500 {object} models.Message
// @Router /change-password [post]
// @Security ApiKeyAuth
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	var req *models.ChangePasswordRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewMessage(err.Error()))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.NewMessage("user not authenticated"))
		return
	}

	cuserr := h.authService.ChangePassword(userID.(int), req)
	if cuserr != nil {
		c.JSON(cuserr.HTTPCode, models.NewMessage(cuserr.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password changed successfully"})
}

// CheckUsernameExists checks if a username exists.
// @Summary Check if a username exists
// @Description Check if a username exists
// @Tags auth
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Param username body models.CheckUsernameRequest true "Change Password Request"
// @Param username formData string false "Username"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /check-username [post]
// @Security ApiKeyAuth
func (h *AuthHandler) CheckUsernameExists(c *gin.Context) {
	var req *models.CheckUsernameRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cuserr := h.authService.CheckUsernameExists(req.Username)
	if cuserr != nil {
		c.JSON(cuserr.HTTPCode, models.NewMessage(cuserr.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "username is available"})
}
