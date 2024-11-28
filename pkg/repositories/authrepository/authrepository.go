package authrepository

import (
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/interfaces/authinterface"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/models/authmodels"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/utils/customerror"
)

// AuthRepository provides methods to interact with the authentication service.
type AuthRepository struct {
	service authinterface.AuthInterface
}

// NewAuthRepository creates a new instance of AuthRepository.
// Parameters:
//   - service: implementation of AuthInterface for handling auth operations
//
// Returns:
//   - *AuthRepository: new repository instance
func NewAuthRepository(service authinterface.AuthInterface) *AuthRepository {
	return &AuthRepository{service: service}
}

// CreateUser delegates user creation to the underlying service.
// Parameters:
//   - username: unique identifier for the user
//   - email: user's email address
//   - password: user's password credential
//
// Returns:
//   - *customerror.CustomError: nil if successful, error details if failed
func (r *AuthRepository) CreateUser(username string, email string, password string) *customerror.CustomError {
	return r.service.CreateUser(username, email, password)
}

// GetUserByEmail delegates email lookup to the underlying service.
// Parameters:
//   - email: string containing the email address to search
//
// Returns:
//   - *User: user information if found, nil if no matching user
//   - error: nil if lookup successful, otherwise contains error details
func (r *AuthRepository) GetUserByEmail(email string) (*authmodels.User, *customerror.CustomError) {
	return r.service.GetUserByEmail(email)
}

// GetUserByID retrieves a user by their ID using the underlying service.
func (r *AuthRepository) GetUserByID(id int) (*authmodels.User, *customerror.CustomError) {
	return r.service.GetUserByID(id)
}

// GetUserByUsername retrieves a user by their username using the underlying service.
func (r *AuthRepository) GetUserByUsername(username string) (*authmodels.User, *customerror.CustomError) {
	return r.service.GetUserByUsername(username)
}

// CheckUsernameExists checks if a username exists using the underlying service.
// Parameters:
//   - username: string containing the username to check
//
// Returns:
//   - bool:
//   - true: username already exists and is not available for use
//   - false: username doesn't exist and is available for registration
//   - error: nil if check successful, otherwise contains error details
func (r *AuthRepository) CheckUsernameExists(username string) (bool, *customerror.CustomError) {
	return r.service.CheckUsernameExists(username)
}

// UpdatePassword updates a user's password using the underlying service.
func (r *AuthRepository) UpdatePassword(userID int, passwordHash string) *customerror.CustomError {
	return r.service.UpdatePassword(userID, passwordHash)
}
